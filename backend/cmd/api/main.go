package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/maxzhirnov/formease/config"
	"github.com/maxzhirnov/formease/internal/handlers"
	"github.com/maxzhirnov/formease/internal/middleware"
	"github.com/maxzhirnov/formease/internal/repository"
	"github.com/maxzhirnov/formease/internal/service"
	"github.com/maxzhirnov/formease/internal/storage"
	"github.com/maxzhirnov/formease/internal/utils"
	"github.com/maxzhirnov/formease/pkg/database"
)

func setupRoutes(router *gin.Engine,
	formHandler *handlers.FormHandler,
	authHandler *handlers.AuthHandler,
	healthHandler *handlers.HealthHandler,
	gptHandler *handlers.GPTHandler,
	imageHandler *handlers.ImageHandler,
	submissionHandler *handlers.SubmissionHandler,
	jwtUtil *utils.JWTUtil) {
	// Public health check routes
	router.GET("/ping", healthHandler.Ping)
	router.GET("/health", healthHandler.HealthCheck)

	api := router.Group("/api/v1")
	{
		// Auth routes (public)
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.Refresh)
		}

		// Public form routes (read-only access)
		forms := api.Group("/forms")
		{
			forms.GET("/:id", formHandler.GetForm) // Get a specific form
		}

		submission := api.Group("/submissions")
		{
			submission.POST("", submissionHandler.CreateSubmission)
		}

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(jwtUtil))
		{
			// User profile routes
			protected.GET("/profile", authHandler.GetProfile)
			// Image upload routes
			protected.POST("/image-upload", imageHandler.UploadImage)
			protected.GET("/images", imageHandler.GetUserImages)
			protected.DELETE("/images/:id", imageHandler.DeleteImage)

			// User's personal forms routes
			userForms := protected.Group("/my-forms")
			{
				userForms.GET("/:id/toggle-draft", formHandler.ToggleDraftStatus)
				userForms.GET("", formHandler.ListForms)         // List user's forms
				userForms.GET("/:id", formHandler.GetForm)       // Get user's specific form
				userForms.POST("", formHandler.CreateForm)       // Create new form
				userForms.PUT("/:id", formHandler.UpdateForm)    // Update user's form
				userForms.DELETE("/:id", formHandler.DeleteForm) // Delete user's form
				userForms.POST("/generate-form", gptHandler.GenerateForm)
			}
		}
	}
}

func setupStaticFileServing(router *gin.Engine, fileStorage storage.FileStorage) {
	router.Use(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/uploads/") {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
		}
		c.Next()
	})

	// Используем кастомный обработчик
	router.GET("/uploads/*filepath", func(c *gin.Context) {
		filepath := c.Param("filepath")

		// Проверяем существование файла через storage
		fullPath := fileStorage.GetFullPath(filepath)

		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			c.JSON(404, gin.H{"error": "File not found"})
			return
		}

		// Отдаем файл
		c.File(fullPath)
	})
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Initialize configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize MongoDB connection
	client, err := database.NewMongoConnection(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Ensure MongoDB disconnection
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Printf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	// Get database instance
	db := client.Database(cfg.MongoDB)

	// Initialize repositories
	formRepo := repository.NewFormRepository(db)
	userRepo := repository.NewUserRepository(db)
	err = userRepo.EnsureIndexes(context.Background())
	if err != nil {
		log.Fatalf("Failed to ensure indexes: %v", err)
	}
	imageRepo := repository.NewMongoImageRepository(db)
	submissionRepo := repository.NewSubmissionRepository(db)

	fileStorageConfig := storage.FileStorageConfig{
		UploadDir:       "uploads", // Не используется для S3
		BaseURL:         "https://storage.yandexcloud.net",
		Endpoint:        "https://storage.yandexcloud.net",
		Region:          "ru-central1",
		AccessKeyID:     os.Getenv("YANDEX_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("YANDEX_SECRET_ACCESS_KEY"),
		BucketName:      "formease",
	}

	fileStorage, err := storage.NewYandexS3Storage(fileStorageConfig)
	if err != nil {
		log.Fatalf("Failed to create S3 storage: %v", err)
	}
	// fileStorage := storage.NewLocalFileStorage(fileStorageConfig)

	// Initialize JWT utility
	jwtUtil := utils.NewJWTUtil(cfg.AuthSecret)

	// Initialize services
	formService := service.NewFormService(formRepo)
	userService := service.NewUserService(cfg, userRepo, jwtUtil)
	gptService := service.NewYandexGPTService("AQVN3j7OW3-zdGmDl4p5nr8D7MHizPCs9tHd0IqG", "b1gakioh5lutqcssd8ph")
	imageService := service.NewImageService(imageRepo, fileStorage)
	submissionService := service.NewSubmissionService(submissionRepo)

	// Initialize handlers
	formHandler := handlers.NewFormHandler(formService)
	authHandler := handlers.NewAuthHandler(userService)
	healthHandler := handlers.NewHealthHandler(client)
	gptHandler := handlers.NewGPTHandler(formService, gptService)
	imageHandler := handlers.NewImageHandler(imageService)
	submissionHandler := handlers.NewSubmissionHandler(submissionService)

	// Set up Gin router
	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	// Передаем fileStorage в setupStaticFileServing
	setupStaticFileServing(router, fileStorage)

	// Routes
	setupRoutes(router, formHandler, authHandler, healthHandler, gptHandler, imageHandler, submissionHandler, jwtUtil)

	// Create server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

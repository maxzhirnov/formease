package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/maxzhirnov/formease/config"
	"github.com/maxzhirnov/formease/internal/handlers"
	"github.com/maxzhirnov/formease/internal/middleware"
	"github.com/maxzhirnov/formease/internal/repository"
	"github.com/maxzhirnov/formease/internal/service"
	"github.com/maxzhirnov/formease/internal/utils"
	"github.com/maxzhirnov/formease/pkg/database"
)

func setupRoutes(router *gin.Engine, formHandler *handlers.FormHandler, authHandler *handlers.AuthHandler, healthHandler *handlers.HealthHandler, jwtUtil *utils.JWTUtil) {
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

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(jwtUtil))
		{
			// User profile routes
			protected.GET("/profile", authHandler.GetProfile)

			// User's personal forms routes
			userForms := protected.Group("/my-forms")
			{
				userForms.GET("", formHandler.ListForms)         // List user's forms
				userForms.GET("/:id", formHandler.GetForm)       // Get user's specific form
				userForms.POST("", formHandler.CreateForm)       // Create new form
				userForms.PUT("/:id", formHandler.UpdateForm)    // Update user's form
				userForms.DELETE("/:id", formHandler.DeleteForm) // Delete user's form
			}
		}
	}
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

	// Initialize JWT utility
	jwtUtil := utils.NewJWTUtil(cfg.AuthSecret)

	// Initialize services
	formService := service.NewFormService(formRepo)
	userService := service.NewUserService(cfg, userRepo, jwtUtil)

	// Initialize handlers
	formHandler := handlers.NewFormHandler(formService)
	authHandler := handlers.NewAuthHandler(userService)
	healthHandler := handlers.NewHealthHandler(client)

	// Set up Gin router
	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.LoggingMiddleware())
	router.Use(middleware.CORSMiddleware())

	// Routes
	setupRoutes(router, formHandler, authHandler, healthHandler, jwtUtil)

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

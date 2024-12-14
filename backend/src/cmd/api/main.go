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
	"github.com/maxzhirnov/formease/internal/repository"
	"github.com/maxzhirnov/formease/internal/service"
	"github.com/maxzhirnov/formease/pkg/database"
)

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

	// Initialize database connection
	db, err := database.NewPostgresConnection(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	formRepo := repository.NewFormRepository(db)
	questionRepo := repository.NewQuestionRepository(db)

	// Initialize services
	formService := service.NewFormService(formRepo, questionRepo)
	submissionService := service.NewSubmissionService(formRepo, questionRepo)

	// Initialize handlers
	formHandler := handlers.NewFormHandler(formService)
	submissionHandler := handlers.NewSubmissionHandler(submissionService)

	// Set up Gin router
	router := gin.Default()

	// Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Routes
	v1 := router.Group("/api/v1")
	{
		forms := v1.Group("/forms")
		{
			forms.GET("/:id", formHandler.GetForm)
			forms.GET("/:id/questions", formHandler.GetFormQuestions)
			forms.POST("/:id/submit", submissionHandler.SubmitForm)
			forms.GET("/:id/thank-you", formHandler.GetThankYouMessage)
			forms.POST("", formHandler.CreateForm)
			forms.PUT("/:id", formHandler.UpdateForm)
			forms.DELETE("/:id", formHandler.DeleteForm)
		}
	}

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

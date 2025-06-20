package main

import (
	"context"
	"errors"
	"event-booking/db"
	"event-booking/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found. Using system environment variables.")
	}

	// Set Gin mode based on environment
	if os.Getenv("ENVIRONMENT") == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize sqlite db connection
	db.InitSqliteDB()
	defer db.CloseSqlite3DB()

	// Setup Gin server
	server := gin.Default()
	routes.RegisterRoutes(server)

	// Create HTTP server with timeouts
	srv := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      server,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Channel to listen for errors from server goroutine
	serverErr := make(chan error, 1)

	// WaitGroup to ensure we don't exit before cleanup
	var wg sync.WaitGroup
	wg.Add(1)

	// Run server in a goroutine
	go func() {
		defer wg.Done()
		log.Printf("Server starting on port %s", os.Getenv("PORT"))
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	// Graceful shutdown channel
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Wait for either an error or shutdown signal
	select {
	case err := <-serverErr:
		log.Fatalf("Server error: %v", err)
	case <-quit:
		log.Println("Shutting down server...")

		// Create shutdown context with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Shutdown HTTP server
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP server shutdown error: %v", err)
		}

		// Wait for goroutines to finish
		wg.Wait()

		// DB will be closed via `defer db.CloseDB()`
		log.Println("Server exited properly")
	}
}

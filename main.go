package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"git-example/internal/routes"
	"github.com/gin-gonic/gin"
)

const timeout = 5 * time.Second

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	r := gin.Default()

	routes.SetupRoutes(r)

	srv := http.Server{
		Addr:    ":3031",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("error starting server with the following: %v\n", err)
		}
	}()

	<-ctx.Done()
	cancel()
	log.Println("Shutting down server...")

	// set 5 seconds of timeout to force shutdown
	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	if err := srv.Shutdown(timeoutCtx); err != nil {
		log.Fatalf("Server forced shutdown after timeout with err: %v\n", err)
	}

	log.Println("Server shut down...")
}

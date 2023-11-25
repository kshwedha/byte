package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kshwedha/core/common/config"
	"github.com/kshwedha/core/src/api"
)

func Serve() {
	config := config.GetConfig()
	serverPort := config.GetInt("server.port")
	serverHost := config.GetString("server.host")

	r := api.Routes()
	http.Handle("/", r)
	fmt.Printf("Server is hosted on host:port %s:%d", serverHost, serverPort)

	// Create an HTTP server with the desired configuration
	server := &http.Server{
		Addr:    serverHost + ":" + fmt.Sprint(serverPort), // Port to listen on
		Handler: r,
	}

	// Start the server in a separate goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %s", err)
		}
	}()

	// Create a channel to receive OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Wait for an interrupt signal (e.g., Ctrl+C)
	<-stop

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shut down the server gracefully
	fmt.Println("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown error: %s", err)
	}

	fmt.Println("Server gracefully stopped")
}

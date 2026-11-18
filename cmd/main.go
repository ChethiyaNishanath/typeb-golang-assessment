package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ChethiyaNishanath/typeb-golang-assessment/internal/greeting"
	"github.com/ChethiyaNishanath/typeb-golang-assessment/internal/handlers"
	"github.com/ChethiyaNishanath/typeb-golang-assessment/internal/server"
)

func main() {

	greeter := greeting.NewService()
	helloWorldHandler := handlers.NewHelloWorldHandler(greeter)
	server.RegisterRoutes(helloWorldHandler)

	srv := server.NewServer(":8080")
	errChan := srv.Start()

	// Graceful shutdown handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-errChan:
		if err != nil {
			fmt.Println("Error starting server:", err)
			os.Exit(1)
		}
	case <-sigChan:
		fmt.Println("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println("Error during shutdown:", err)
		}
		fmt.Println("Server stopped")
	}

}

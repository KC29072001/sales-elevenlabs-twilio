// package main

// import (
// 	"context"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"caller/internal/config"
// 	"caller/internal/server"
// )

// func main() {
// 	if err := run(); err != nil {
// 		log.Printf("error: %v\n", err)
// 		os.Exit(1)
// 	}
// }

// func run() error {
// 	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
// 	defer cancel()

// 	cfg, err := config.Load()
// 	if err != nil {
// 		return err
// 	}

// 	srv, err := server.New(cfg)
// 	if err != nil {
// 		return err
// 	}

// 	go func() {
// 		if err := srv.Start(); err != nil {
// 			log.Printf("server error: %v\n", err)
// 		}
// 	}()

// 	<-ctx.Done()

// 	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer shutdownCancel()

// 	if err := srv.Shutdown(shutdownCtx); err != nil {
// 		return err
// 	}

// 	return nil
// }



package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv" // Import the godotenv package

	"caller/internal/config"
	"caller/internal/server"
)

func main() {
	if err := run(); err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
		// You might choose to return the error here, or handle it in some other way
		// For example, you could proceed with default configuration values.
		// For this example, we'll return the error to stop execution.
		return err 
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		return err
	}

	srv, err := server.New(cfg)
	if err != nil {
		return err
	}

	go func() {
		if err := srv.Start(); err != nil {
			log.Printf("server error: %v\n", err)
		}
	}()

	<-ctx.Done()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}

	return nil
}

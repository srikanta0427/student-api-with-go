package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/srikanta0427/student-api/internal/config"
	"github.com/srikanta0427/student-api/internal/http/student"
)

func main() {
	// load config
	cfg := config.MustLoad()
	slog.Info("Config loaded successfully", slog.String("env", cfg.Env), slog.String("storagePath", cfg.StoragePath), slog.String("httpServerHost", cfg.HttpServer.Host), slog.String("httpServerPort", cfg.HttpServer.Port))

	// setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /api/students", student.New())

	// setup server
	server := http.Server{
		Addr:    cfg.HttpServer.Host + ":" + cfg.HttpServer.Port,
		Handler: router,
	}

	fmt.Println("Server started on port", cfg.HttpServer.Port)

	// for graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("error starting server: %v", err.Error())
		}
	}()

	<-done

	slog.Info("Server stopped gracefully")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)

	defer cancel()

	err := server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("Server Shutdown ")
}

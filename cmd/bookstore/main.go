package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ryoeuyo/test_go/internal/application/service"
	"github.com/ryoeuyo/test_go/internal/config"
	"github.com/ryoeuyo/test_go/internal/infrastructure/http/handlers/checks/ping"
	"github.com/ryoeuyo/test_go/internal/infrastructure/http/handlers/crud"
	"github.com/ryoeuyo/test_go/internal/infrastructure/http/handlers/middleware"
	"github.com/ryoeuyo/test_go/internal/infrastructure/repository/postgres"
	"github.com/ryoeuyo/test_go/internal/shared/logger"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	logger := logger.Setup(cfg.Env)

	logger.Debug("config is initalized", slog.Any("environment", cfg.Env))

	conn := postgres.MustConnect(ctx, cfg.Database)
	defer conn.Close(ctx)

	logger.Debug("repository is initialized", slog.Any("postgres config", map[string]interface{}{
		"host":     cfg.Database.Host,
		"port":     cfg.Database.Port,
		"database": cfg.Database.Name,
	}))

	repo := postgres.New(conn)

	service := service.Service{
		Repository: repo,
	}

	r := gin.New()
	r.Use(gin.Recovery(), middleware.SlogMiddleware(logger))
	r.GET("/ping", ping.Ping)
	r.GET("/books", crud.AllBooks(ctx, service))
	r.POST("/book", crud.AddBook(ctx, service))
	r.GET("/book/", crud.GetBook(ctx, service))

	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", cfg.HTTPServer.Address, cfg.HTTPServer.Port),
		Handler:        r,
		ReadTimeout:    cfg.HTTPServer.Timeout,
		WriteTimeout:   cfg.HTTPServer.Timeout,
		IdleTimeout:    cfg.HTTPServer.IdleTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := srv.ListenAndServe()
	if err != nil {
		logger.Error("error start server", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// TODO: init metrics
	// TODO: Start application
}

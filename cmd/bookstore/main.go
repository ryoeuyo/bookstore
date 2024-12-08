package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ryoeuyo/bookstore/internal/application/service"
	"github.com/ryoeuyo/bookstore/internal/config"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/checks/health"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/checks/ping"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/crud"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/middleware"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
	"github.com/ryoeuyo/bookstore/internal/shared/logger"
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

	service := &service.BookService{
		Repository: repo,
	}

	baseRouter := gin.New()
	baseRouter.Use(gin.Recovery(), middleware.SlogLogger(logger))

	apiRoute := baseRouter.Group("/api")
	{
		bookRoute := apiRoute.Group("/books")
		bookRoute.GET("/", crud.AllBooks(ctx, service))
		bookRoute.POST("/", crud.AddBook(ctx, service))
		bookRoute.GET("/:id", crud.GetBook(ctx, service))
		bookRoute.DELETE("/:id", crud.DeleteBook(ctx, service))
	}

	healthRoute := baseRouter.Group("/health-check")
	{
		healthRoute.GET("/ping", ping.Ping())
		healthRoute.GET("/check", health.Check(conn))
	}

	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", cfg.HTTPServer.Address, cfg.HTTPServer.Port),
		Handler:        baseRouter,
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

	// TODO: write swagger docs
	// TODO: init metrics
	// TODO: Start application
}

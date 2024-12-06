package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/ryoeuyo/test_go/internal/config"
	"github.com/ryoeuyo/test_go/internal/infrastructure/http/checks/ping"
	"github.com/ryoeuyo/test_go/internal/infrastructure/repository/postgres"
	"github.com/ryoeuyo/test_go/internal/shared/logger"
)

func main() {
	ctx := context.Background()
	cfg := config.MustLoad()
	logger := logger.Setup(cfg.Env)

	logger.Debug("config is initalized", slog.Any("environment", cfg.Env))

	conn, err := postgres.Connect(ctx, cfg.Database)
	if err != nil {
		logger.Error("failed connect to database", slog.String("error", err.Error()))
	}

	logger.Debug("repository is initialized", slog.Any("postgres config", map[string]interface{}{
		"host":     cfg.Database.Host,
		"port":     cfg.Database.Port,
		"database": cfg.Database.Name,
	}))

	repo := postgres.New(conn)

	_, _ = repo.AllBooks(ctx)
	r := gin.Default()
	r.GET("/ping", ping.Ping)
	r.Run(fmt.Sprintf("%s:%d", cfg.HTTPServer.Address, cfg.HTTPServer.Port))

	// TODO: init handlers
	// TODO: init metrics
	// TODO: Start application
}

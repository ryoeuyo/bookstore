package main

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryoeuyo/test_go/internal/config"
	"github.com/ryoeuyo/test_go/internal/shared/logger"
)

func main() {
	cfg := config.MustLoad()

	logger := logger.Setup(cfg.Env)
	logger.Info("Config is initalized", slog.Any("config", cfg))

	// TODO: init repository
	// TODO: create migrations
	// TODO: init router
	// TODO: init handlers
	// TODO: init metrics
	// TODO: Start application
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "alive",
	})
}

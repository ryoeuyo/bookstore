package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func SlogLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		if len(c.Errors) > 0 {
			logger.Error(
				"request errors",
				slog.String("method", c.Request.Method),
				slog.String("request_uri", c.Request.RequestURI),
				slog.Any("header", c.Request.Header),
				slog.Any("errors", c.Errors.Errors()),
			)
			return
		}

		logger.Info(
			"request processed",
			slog.String("method", c.Request.Method),
			slog.String("request_uri", c.Request.RequestURI),
			slog.Any("header", c.Request.Header),
			slog.Any("time", time.Since(start)),
		)
	}
}

package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func SlogMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info(
			"incoming request",
			slog.String("method", c.Request.Method),
			slog.String("request_uri", c.Request.RequestURI),
			slog.String("user_agent", c.Request.UserAgent()),
			slog.Any("header", c.Request.Header),
		)

		start := time.Now()
		c.Next()

		if len(c.Errors) > 0 {
			logger.Error(
				"request errors",
				slog.String("method", c.Request.Method),
				slog.String("request_uri", c.Request.RequestURI),
				slog.Any("errors", c.Errors.Errors()),
			)

			return
		}

		logger.Info(
			"request processed",
			slog.String("method", c.Request.Method),
			slog.String("request_uri", c.Request.RequestURI),
			slog.String("user_agent", c.Request.UserAgent()),
			slog.Any("header", c.Request.Header),
			slog.Any("time", time.Since(start)),
		)
	}
}

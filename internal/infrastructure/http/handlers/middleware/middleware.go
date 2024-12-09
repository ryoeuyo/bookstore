package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/metric"
)

func SlogLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		if len(c.Errors) > 0 {
			logger.Error(
				"request error",
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

func IncRequest(m *metric.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		m.RequestCount.Inc()

		errsCount := float64(len(c.Errors))
		m.ErrorsCount.Add(errsCount)

		c.Next()
	}
}

func ObserveRequest(m *metric.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		defer func(start time.Time) {
			method := c.Request.Method
			elapsed := time.Since(start).Seconds()

			m.RequestDuration.WithLabelValues(method).Observe(elapsed)
		}(start)
	}
}

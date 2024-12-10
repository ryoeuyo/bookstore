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
		defer func() {
			if len(c.Errors) > 0 {
				logger.Error(
					"",
					slog.String("method", c.Request.Method),
					slog.String("request_uri", c.Request.RequestURI),
					slog.Any("errors", c.Errors.Errors()),
				)
				return
			}

			logger.Info(
				"",
				slog.String("method", c.Request.Method),
				slog.String("request_uri", c.Request.RequestURI),
				slog.Any("time", time.Since(start)),
			)
		}()

		c.Next()
	}
}

func IncRequest(m *metric.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			m.RequestCount.Inc()

			errsCount := float64(len(c.Errors))
			m.ErrorsCount.Add(errsCount)
		}()

		c.Next()
	}
}

func ObserveRequest(m *metric.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		defer func(start time.Time) {
			method := c.Request.Method
			elapsed := time.Since(start).Seconds()

			m.RequestDuration.WithLabelValues(method).Observe(elapsed)
		}(start)

		c.Next()
	}
}

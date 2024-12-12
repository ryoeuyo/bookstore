package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ryoeuyo/bookstore/internal/application/service"
	"github.com/ryoeuyo/bookstore/internal/config"
	"github.com/ryoeuyo/bookstore/internal/etc/logger"
	"github.com/ryoeuyo/bookstore/internal/etc/validate"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/checks/health"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/checks/ping"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/crud"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/handlers/middleware"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/http/metric"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
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
	metrics := metric.NewMetrics()
	svc := service.NewBookService(repo)

	valid := validator.New()
	valid.RegisterValidation("notzero", validate.IsNotZero)
	valid.RegisterValidation("notempty", validate.IsNotEmpty)

	handler := crud.NewBookHandler(svc, valid)

	router := gin.New()
	router.Use(
		gin.Recovery(),
		middleware.IncRequest(metrics),
		middleware.ObserveRequest(metrics),
		middleware.SlogLogger(logger),
	)

	apiRouter := router.Group("/api/v1")
	{
		bookRoute := apiRouter.Group("/books")
		bookRoute.GET("/", handler.AllBooks(ctx))
		bookRoute.POST("/", handler.AddBook(ctx))
		bookRoute.GET("/:id", handler.GetBook(ctx))
		bookRoute.DELETE("/:id", handler.DeleteBook(ctx))
		bookRoute.PUT("/", handler.UpdateBook(ctx))
		bookRoute.PATCH("/", handler.UpdateFieldBook(ctx))
	}

	healthRouter := router.Group("/health-check")
	{
		healthRouter.GET("/ping", ping.Ping())
		healthRouter.GET("/check", health.Check(conn))
	}

	srv := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", cfg.HTTPServer.Address, cfg.HTTPServer.Port),
		Handler:        router,
		ReadTimeout:    cfg.HTTPServer.Timeout,
		WriteTimeout:   cfg.HTTPServer.Timeout,
		IdleTimeout:    cfg.HTTPServer.IdleTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err := metric.StartMetricServer(cfg.MetricServer.Address, int(cfg.MetricServer.Port))
		if err != nil {
			logger.Error("error start metric server", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	err := srv.ListenAndServe()
	if err != nil {
		logger.Error("error start server", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// TODO: write swagger docs
}

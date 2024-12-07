package crud

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryoeuyo/test_go/internal/application/service"
)

func AllBooks(ctx context.Context, service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := service.AllBooks(ctx)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"error":  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": books,
		})
	}
}

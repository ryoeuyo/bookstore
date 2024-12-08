package crud

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryoeuyo/bookstore/internal/application/service"
)

func AllBooks(ctx context.Context, service *service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := service.AllBooks(ctx)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, books)
	}
}

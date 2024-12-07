package crud

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryoeuyo/test_go/internal/application/service"
	"github.com/ryoeuyo/test_go/internal/domain/book"
)

func AddBook(ctx context.Context, service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book book.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  err.Error(),
			})
			return
		}

		id, err := service.AddBook(ctx, book)
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
			"message": "successfully added",
			"id":      id,
		})
	}
}

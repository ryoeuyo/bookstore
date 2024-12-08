package crud

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ryoeuyo/bookstore/internal/application/service"
)

func GetBook(ctx context.Context, service *service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := uuid.Parse(idParam)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id is invalid",
			})
			return
		}

		book, err := service.GetBook(ctx, id)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, book)
	}
}

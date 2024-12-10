package crud

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryoeuyo/bookstore/internal/application/service"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
)

func UpdateBook(ctx context.Context, s *service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bookUpdateParams postgres.UpdateBookParams
		if err := c.ShouldBindJSON(&bookUpdateParams); err != nil {
			c.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": ErrDeserialize,
			})
			return
		}

		id, err := s.UpdateBook(ctx, bookUpdateParams)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "updated",
			"id":      id,
		})
	}
}

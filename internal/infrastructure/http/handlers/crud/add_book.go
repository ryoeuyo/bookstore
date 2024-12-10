package crud

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ryoeuyo/bookstore/internal/application/service"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
)

func AddBook(ctx context.Context, s *service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book postgres.AddBookParams
		if err := c.ShouldBindJSON(&book); err != nil {
			c.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": errors.New(ErrDeserialize),
			})
			return
		}

		id, err := s.AddBook(ctx, book)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	}
}

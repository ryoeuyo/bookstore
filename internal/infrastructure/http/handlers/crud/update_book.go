package crud

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

		err := s.Validate.Struct(bookUpdateParams)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				c.JSON(http.StatusInternalServerError, gin.H{
					"errror": err.Error(),
				})
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"errror": err.Error(),
			})
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

func UpdateFieldBook(ctx context.Context, s *service.BookService) gin.HandlerFunc {
	return func(c *gin.Context) {
		field := c.Query("field")
		if field == "" {
			c.Error(errors.New(ErrDeserialize))
			c.JSON(http.StatusBadRequest, gin.H{
				"error": ErrDeserialize,
			})
			return
		}

		var req UpdateFieldBookRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(errors.New(ErrDeserialize))
			c.JSON(http.StatusBadRequest, gin.H{
				"error": ErrDeserialize,
			})
			return
		}

		err := s.Validate.Struct(req)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				c.JSON(http.StatusInternalServerError, gin.H{
					"errror": err.Error(),
				})

				return
			}

			c.JSON(http.StatusBadRequest, gin.H{
				"errror": ErrValidation,
			})
			return
		}

		id, err := s.UpdateFieldBook(ctx, req.ID, field, req.Value)
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

package crud

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ryoeuyo/bookstore/internal/infrastructure/repository/postgres"
)

func (h *BookHandler) UpdateBook(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bookUpdateParams postgres.UpdateBookParams
		if err := c.ShouldBindJSON(&bookUpdateParams); err != nil {
			c.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": ErrInvalidJSONRequest,
			})
			return
		}

		err := h.Valid.Struct(bookUpdateParams)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				c.Error(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			c.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		id, err := h.Svc.UpdateBook(ctx, bookUpdateParams)
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

func (h *BookHandler) UpdateFieldBook(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		field := c.Query("field")
		if field == "" {
			c.Error(errors.New(ErrInvalidField))
			c.JSON(http.StatusBadRequest, gin.H{
				"error": c.Errors.Errors(),
			})
			return
		}

		var req UpdateFieldBookRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.Error(errors.New(ErrInvalidField))
			c.JSON(http.StatusBadRequest, gin.H{
				"error": c.Errors.Errors(),
			})
			return
		}

		err := h.Valid.Struct(req)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				c.Error(errors.New(ErrValidation))
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": c.Errors.Errors(),
				})

				return
			}

			c.Error(errors.New(ErrInvalidJSONRequest))
			c.JSON(http.StatusBadRequest, gin.H{
				"error": c.Errors.Errors(),
			})
			return
		}

		id, err := h.Svc.UpdateFieldBook(ctx, req.ID, field, req.Value)
		if err != nil {
			c.Error(errors.New(ErrInternalServer))
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": c.Errors.Errors(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "updated",
			"id":      id,
		})
	}
}

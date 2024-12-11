package crud

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BookHandler) AllBooks(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		books, err := h.Svc.AllBooks(ctx)
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

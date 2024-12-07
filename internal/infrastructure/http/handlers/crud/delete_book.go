package crud

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ryoeuyo/test_go/internal/application/service"
)

func DeleteBook(ctx context.Context, service service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		idQuery := c.Query("id")
		id, err := uuid.Parse(idQuery)
		if err != nil {
			c.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"status": http.StatusBadRequest,
				"error":  "id is invalid",
			})
			return
		}

		id, err = service.DeleteBook(ctx, id)
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
			"message": "book successfully deleted",
			"id":      id,
		})
	}
}

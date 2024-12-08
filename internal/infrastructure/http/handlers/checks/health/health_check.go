package health

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func Check(connDB *pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := connDB.Ping(context.Background()); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"warning": "failed ping database!",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

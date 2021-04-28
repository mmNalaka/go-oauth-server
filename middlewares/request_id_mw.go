package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid, err := uuid.NewV4()

		if err != nil {
			log.Fatalf("failed to generate UUID: %v", err)
		}

		c.Writer.Header().Set("X-Request-Id", uuid.String())
		c.Next()
	}
}

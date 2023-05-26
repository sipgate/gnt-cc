package middleware

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		lastErr := c.Errors.Last()

		if lastErr == nil {
			return
		}

		c.AbortWithStatusJSON(500, lastErr)
	}
}

package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Validate[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload T

		// Bind JSON request body to struct
		if err := c.ShouldBind(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid format"})
			c.Abort()
			return
		}

		// Validate struct fields
		if err := validate.Struct(payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Store validated payload in context
		c.Set("payload", payload)
		c.Next()
	}
}

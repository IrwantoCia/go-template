package logs

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GinMiddlewareLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Generate a unique trace ID
		traceID := uuid.New().String()

		// Attach trace ID to context
		c.Set("traceID", traceID)

		// Process request
		c.Next()

		// Log request details with trace ID
		Logger.Info("HTTP Request",
			zap.String("trace_id", traceID),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)),
			zap.String("client_ip", c.ClientIP()),
		)
	}
}

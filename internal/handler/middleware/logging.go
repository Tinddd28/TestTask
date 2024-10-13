package middleware

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"time"
)

func LoggingMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()

		logger.Info("Request",
			"status", statusCode,
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"duration", duration)
	}
}

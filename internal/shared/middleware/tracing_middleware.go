package middleware

import (
	"github.com/gin-gonic/gin"
	"homework/internal/shared/logger"
	"homework/internal/shared/ulid"
)

func NewTracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := logger.WithData(c.Request.Context(), "request_id", ulid.NewULID().String())
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

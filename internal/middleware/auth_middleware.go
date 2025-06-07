package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/still-breath/go-pos-backend.git/internal/handler" // Ganti path
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := handler.Store.Get(c.Request, "pos-session")

		userID, ok := session.Values["user_id"]
		if !ok || userID == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

package middlewares

import (
	"net/http"
	"rest-api/internal/domain"
	"rest-api/internal/infra/services"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService *services.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format. Use 'Bearer TOKEN'"})
			return
		}

		tokenStr := parts[1]

		user, err := jwtService.ValidateCode(tokenStr)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func RoleMiddleware(requiredRole domain.UserRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, exists := c.Get("user")

		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			return
		}

		user := val.(services.JwtData)

		if user.UserRole != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You do not have permission"})
			return
		}

		c.Next()
	}
}

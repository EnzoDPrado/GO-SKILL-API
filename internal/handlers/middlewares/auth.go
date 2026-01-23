package middlewares

import (
	"fmt"
	"net/http"
	"rest-api/internal/infra/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService *services.JwtService) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr, err := c.Cookie("Auth")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("Expired token")})
			return
		}

		userID, err := jwtService.ValidateCode(tokenStr)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Errorf("Unauthorized")})
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}

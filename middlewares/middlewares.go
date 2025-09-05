package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var secretKey = os.Getenv("JWT_SECRET")
		authHeader := c.GetHeader("Authorization")
		splittedHeader := strings.Split(authHeader, "Bearer ")
		if len(splittedHeader) != 2 {
			c.AbortWithStatusJSON(401, gin.H{"error": "Malformed Token"})
			return
		}

		jwtToken := splittedHeader[1]
		token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}

		// ✅ extract user_id from claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if userID, exists := claims["user_id"]; exists {
				c.Set("user_id", userID) // store in context
			}
		}

		c.Next()
	}
}

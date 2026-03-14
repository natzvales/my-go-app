package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/natz/go-lib-app/internal/shared/contracts"
)

const userKey = "user"

func AuthMiddleware(secret string) gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		claims := token.Claims.(*auth.Claims)

		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)

		c.Next()
	}
}

func GetUser(c *gin.Context) contracts.User {

	user, _ := c.Get(userKey)

	return user.(contracts.User)
}

// ParseToken is a placeholder implementation. Replace with your actual token parsing logic.
// func ParseToken(token string) (string, error) {
// 	// Example: return token as userID if not empty
// 	if token == "" {
// 		return "", errors.New("empty token")
// 	}
// 	return token, nil
// }

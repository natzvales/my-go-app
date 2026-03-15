package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/natz/go-lib-app/internal/response"
	"github.com/natz/go-lib-app/internal/shared/contracts"
	jwtutil "github.com/natz/go-lib-app/internal/utils/jwt"
)

var userService contracts.UserService

func SetUserService(service contracts.UserService) {
	userService = service
}

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			// c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			response.Unauthorized(c, "Unauthorized")
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		user, err := jwtutil.ValidateToken(tokenString, userService)
		if err != nil {
			// c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token or user not found"})
			response.Unauthorized(c, "Invalid token or user not found")
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}

func GetUser(c *gin.Context) contracts.User {

	user, _ := c.Get("user")

	return user.(contracts.User)
}

func GetRole(c *gin.Context) string {

	user := GetUser(c)

	return user.Role
}

// ParseToken is a placeholder implementation. Replace with your actual token parsing logic.
// func ParseToken(token string) (string, error) {
// 	// Example: return token as userID if not empty
// 	if token == "" {
// 		return "", errors.New("empty token")
// 	}
// 	return token, nil
// }

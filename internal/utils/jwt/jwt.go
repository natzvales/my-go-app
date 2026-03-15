package jwtutil

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	config "github.com/natz/go-lib-app/internal/config"
	"github.com/natz/go-lib-app/internal/shared/contracts"
)

var secret = []byte(config.LoadConfig().JWTSecret)

func GenerateToken(userID uuid.UUID, role string, email string) (string, error) {

	claims := contracts.Claims{
		UserID: userID,
		Role:   role,
		Email:  email, // Add this if you want to include email in the claims
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret)
}

func ValidateToken(tokenStr string, service contracts.UserService) (contracts.User, error) {

	// Get the secret key from your configuration
	token, err := jwt.ParseWithClaims(tokenStr, &contracts.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return contracts.User{}, err
	}

	claims := token.Claims.(*contracts.Claims)

	// Check if the token is expired
	if claims.ExpiresAt != nil && time.Now().After(claims.ExpiresAt.Time) {
		return contracts.User{}, errors.New("token expired")
	}

	user, err := service.GetUser(claims.UserID)

	if err != nil {
		return contracts.User{}, err
	}

	return user, nil
}

func ParseToken(tokenStr string) contracts.User {

	token, err := jwt.ParseWithClaims(tokenStr, &contracts.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		// return contracts.User{}, err
	}

	claims := token.Claims.(*contracts.Claims)

	//Querry if the user exists in the database using claims.UserID, if not return an error

	user := contracts.User{
		ID:    claims.UserID,
		Role:  claims.Role,
		Email: claims.Email, // Add this if it's in your JWT claims
	}
	return user

}

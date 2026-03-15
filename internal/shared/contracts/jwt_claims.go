package contracts

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	Email  string    `json:"email"` // Add this if you want to include email in the claims
	jwt.RegisteredClaims
}

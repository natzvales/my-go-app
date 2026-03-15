package contracts

import "github.com/google/uuid"

// import "github.com/golang-jwt/jwt/v5"

type User struct {
	ID    uuid.UUID
	Email string
	Role  string
}

type UserService interface {
	GetUser(id uuid.UUID) (User, error)
}

// type Claims struct {
// 	UserID string `json:"user_id"`
// 	Role   string `json:"role"`
// 	jwt.RegisteredClaims
// }

package auth

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	Name     string    `json:"name"`
	Email    string    `gorm:"uniqueIndex" json:"email"`
	Password string    `json:"-"`
	Role     string    `json:"role"`
}

package auth

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey"` // UUID primary key
	Name     string    `gorm:"not null"`
	Email    string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`              // hashed password
	Role     string    `gorm:"default:user;not null"` // user role: admin, librarian, member
}

// BeforeCreate GORM hook: auto-generate UUID before inserting record
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

// type User struct {
// 	ID       uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
// 	Name     string    `json:"name"`
// 	Email    string    `gorm:"uniqueIndex" json:"email"`
// 	Password string    `json:"-"`
// 	Role     string    `json:"role"`
// }

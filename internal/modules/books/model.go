package books

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
}

func (u *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return
}

// AutoMigrate function
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}

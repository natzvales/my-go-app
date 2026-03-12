package books

import "gorm.io/gorm"

type Book struct {
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// AutoMigrate function
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}

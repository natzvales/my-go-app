package container

import "gorm.io/gorm"

type Container struct {
	DB *gorm.DB
}

func NewContainer(db *gorm.DB) *Container {
	return &Container{
		DB: db,
	}
}

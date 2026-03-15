package container

import (
	"github.com/natz/go-lib-app/internal/config"
	"gorm.io/gorm"
)

type Container struct {
	DB     *gorm.DB
	Config *config.Config
}

func NewContainer(db *gorm.DB, cfg *config.Config) *Container {
	return &Container{
		DB:     db,
		Config: cfg,
	}
}

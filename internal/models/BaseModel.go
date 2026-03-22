package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return
}

// BeforeUpdate GORM hook: update UpdatedAt timestamp
func (base *BaseModel) BeforeUpdate(tx *gorm.DB) (err error) {
	base.UpdatedAt = time.Now()
	return
}

// Soft delete: set DeletedAt timestamp instead of hard deleting record
func (base *BaseModel) BeforeDelete(tx *gorm.DB) (err error) {
	now := time.Now()
	tx.Statement.SetColumn("DeletedAt", &now)
	return
}

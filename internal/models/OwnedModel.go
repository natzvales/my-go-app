package models

import "github.com/google/uuid"

type OwnedModel struct {
	UserID uuid.UUID `gorm:"type:uuid;index" json:"user_id"`
}

package models

import (
	"time"
)

type ArchivableModel struct {
	BaseModel
	ArchivedAt *time.Time `gorm:"index" json:"archived_at,omitempty"`
}

// Archive sets the ArchivedAt timestamp to mark the record as archived
func (a *ArchivableModel) Archive() {
	now := time.Now()
	a.ArchivedAt = &now
}

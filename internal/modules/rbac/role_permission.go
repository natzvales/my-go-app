package rbac

import "github.com/google/uuid"

type RolePermission struct {
	RoleID       uuid.UUID `json:"role_id" gorm:"type:uuid;not null"`
	PermissionID uuid.UUID `json:"permission_id" gorm:"type:uuid;not null"`
}

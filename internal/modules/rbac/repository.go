package rbac

import "gorm.io/gorm"

type Repository interface {
	RoleHasPermission(roleName string, permission string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) RoleHasPermission(roleName string, permissionName string) (bool, error) {
	var count int64
	err := r.db.Table("roles").
		Joins("JOIN role_permissions ON roles.id = role_permissions.role_id").
		Joins("JOIN permissions ON role_permissions.permission_id = permissions.id").
		Where("roles.name = ? AND permissions.name = ?", roleName, permissionName).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

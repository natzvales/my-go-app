package rbac

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) HasPermission(roleName string, permissionName string) (bool, error) {
	return s.repo.RoleHasPermission(roleName, permissionName)
}

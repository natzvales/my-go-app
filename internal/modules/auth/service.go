package auth

import (
	"github.com/google/uuid"
	appErrors "github.com/natz/go-lib-app/internal/errors"
	"github.com/natz/go-lib-app/internal/shared/contracts"
	jwtutil "github.com/natz/go-lib-app/internal/utils/jwt"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(dto RegisterDTO) (User, error) {

	hash, err := HashPassword(dto.Password)
	if err != nil {
		return User{}, err
	}

	user := User{
		ID:       uuid.New(),
		Email:    dto.Email,
		Name:     dto.Name,
		Password: hash,
	}

	err = s.repo.Create(&user)

	return user, err
}

func (s *Service) Login(dto LoginDTO) (string, error) {

	user, err := s.repo.FindByEmail(dto.Email)
	if err != nil {
		return "", appErrors.New(401, "invalid credentials")
	}

	if !VerifyPassword(dto.Password, user.Password) {
		return "", appErrors.New(401, "invalid credentials")
	}

	// Convert user.ID from string to uint
	// id, err := strconv.ParseUint(user.ID, 10, 64)
	// if err != nil {
	// 	return "", appErrors.New(500, "invalid user ID")
	// }
	return jwtutil.GenerateToken(user.ID, user.Role, user.Email)
}

func (s *Service) GetUser(id uuid.UUID) (contracts.User, error) {
	user, err := s.repo.FindByID(id)

	if err != nil {
		return contracts.User{}, err
	}

	return contracts.User{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}, nil
}

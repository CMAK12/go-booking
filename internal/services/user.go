package service

import (
	"go-booking/internal/models"
	"go-booking/internal/storage"

	"github.com/google/uuid"
)

type userService struct {
	userStorage storage.UserStorage
}

func NewUserService(userStorage storage.UserStorage) UserService {
	return &userService{userStorage: userStorage}
}

func (s *userService) Get(id uuid.UUID) (*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userService) List() ([]*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userService) Create(user *models.User) (*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userService) Update(user *models.User) (*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userService) Delete(id uuid.UUID) error {
	// Implementation here
	return nil
}

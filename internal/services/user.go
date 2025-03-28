package service

import (
	"context"
	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
)

type userService struct {
	userStorage storage.UserStorage
}

func NewUserService(userStorage storage.UserStorage) UserService {
	return &userService{userStorage: userStorage}
}

func (s *userService) Get(ctx context.Context, id string) (*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userService) List(ctx context.Context, filter storage.ListUserFilter) ([]*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userService) Create(ctx context.Context, user *dto.CreateUserRequest) (*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userService) Update(ctx context.Context, user *models.User) (*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	// Implementation here
	return nil
}

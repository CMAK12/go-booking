package service

import (
	"context"
	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
	"log"
)

type userService struct {
	userStorage storage.UserStorage
}

func NewUserService(userStorage storage.UserStorage) UserService {
	return &userService{userStorage: userStorage}
}

func (s *userService) List(ctx context.Context, filter dto.ListUserFilter) ([]models.User, int64, error) {
	users, count, err := s.userStorage.List(ctx, filter)
	if err != nil {
		log.Println("error listing users", err)
		return nil, 0, err
	}
	return users, count, nil
}

func (s *userService) Create(ctx context.Context, dto dto.CreateUserRequest) (models.User, error) {
	user := models.NewUser(dto.Email, dto.Name, dto.Password)

	createdUser, err := s.userStorage.Create(ctx, user)
	if err != nil {
		log.Println("error creating user", err)
		return models.User{}, err
	}

	return createdUser, nil
}

func (s *userService) Update(ctx context.Context, id string, dto dto.UpdateUserRequest) (models.User, error) {
	user := models.NewUpdateUser(id, dto.Name, dto.Email, dto.Password, dto.Role)

	updatedUser, err := s.userStorage.Update(ctx, id, user)
	if err != nil {
		log.Println("error updating user", err)
		return models.User{}, err
	}

	return updatedUser, nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	err := s.userStorage.Delete(ctx, id)
	if err != nil {
		log.Println("error deleting user", err)
		return err
	}
	return nil
}

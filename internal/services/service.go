package service

import (
	"context"

	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
)

type (
	UserService interface {
		Get(ctx context.Context, id string) (*models.User, error)
		List(ctx context.Context, filter storage.ListUserFilter) ([]*models.User, error)
		Create(ctx context.Context, user *dto.CreateUserRequest) (*models.User, error)
		Update(ctx context.Context, user *models.User) (*models.User, error)
		Delete(ctx context.Context, id string) error
	}

	ReservationService interface {
		Get(ctx context.Context, id string) (*models.Reservation, error)
		List(ctx context.Context, filter storage.ListReservationFilter) ([]*models.Reservation, error)
		Create(ctx context.Context, reservation *dto.CreateReservationRequest) (*models.Reservation, error)
		Update(ctx context.Context, reservation *models.Reservation) (*models.Reservation, error)
		Delete(ctx context.Context, id string) error
	}
)

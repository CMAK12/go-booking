package storage

import (
	"context"
	"go-booking/internal/models"
)

const (
	userTable = "users"
	reservationTable = "reservations"
)

type (
	UserStorage interface {
		Get(ctx context.Context, id string) (*models.User, error)
		List(ctx context.Context, filter ListUserFilter) ([]*models.User, error)
		Create(ctx context.Context, user *models.User) (*models.User, error)
		Update(ctx context.Context, user *models.User) (*models.User, error)
		Delete(ctx context.Context, id string) error
	}

	ReservationStorage interface {
		Get(ctx context.Context, id string) (*models.Reservation, error)
		List(ctx context.Context, filter ListReservationFilter) ([]*models.Reservation, error)
		Create(ctx context.Context, reservation *models.Reservation) (*models.Reservation, error)
		Update(ctx context.Context, reservation *models.Reservation) (*models.Reservation, error)
		Delete(ctx context.Context, id string) error
	}
)

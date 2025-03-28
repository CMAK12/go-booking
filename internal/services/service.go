package service

import (
	"go-booking/internal/models"

	"github.com/google/uuid"
)

type (
	UserService interface {
		Get(id uuid.UUID) (*models.User, error)
		List() ([]*models.User, error)
		Create(user *models.User) (*models.User, error)
		Update(user *models.User) (*models.User, error)
		Delete(id uuid.UUID) error
	}

	ReservationService interface {
		Get(id uuid.UUID) (*models.Reservation, error)
		List() ([]*models.Reservation, error)
		Create(reservation *models.Reservation) (*models.Reservation, error)
		Update(reservation *models.Reservation) (*models.Reservation, error)
		Delete(id uuid.UUID) error
	}
)

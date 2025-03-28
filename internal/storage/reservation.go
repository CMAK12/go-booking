package storage

import (
	"context"

	"go-booking/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	reservationStorage struct {
		db *pgxpool.Pool
	}

	ListReservationFilter struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Date      string `json:"date"`
		City 		 	string `json:"city"`
	}
)

func NewReservationStorage(db *pgxpool.Pool) ReservationStorage {
	return &reservationStorage{db: db}
}

func (s *reservationStorage) Get(ctx context.Context, id string) (*models.Reservation, error) {
	// Implementation here
	return nil, nil
}

func (s *reservationStorage) List(ctx context.Context, filter ListReservationFilter) ([]*models.Reservation, error) {
	// Implementation here
	return nil, nil
}

func (s *reservationStorage) Create(ctx context.Context, reservation *models.Reservation) (*models.Reservation, error) {
	// Implementation here
	return nil, nil
}

func (s *reservationStorage) Update(ctx context.Context, reservation *models.Reservation) (*models.Reservation, error) {
	// Implementation here
	return nil, nil
}

func (s *reservationStorage) Delete(ctx context.Context, id string) error {
	// Implementation here
	return nil
}

package storage

import (
	"database/sql"

	"go-booking/internal/models"

	"github.com/google/uuid"
)

type reservationStorage struct {
	db *sql.DB
}

func NewReservationStorage(db *sql.DB) ReservationStorage {
	return &reservationStorage{db: db}
}

func (s *reservationStorage) Get(id uuid.UUID) (*models.Reservation, error) {
	// Implementation here
	return nil, nil
}

func (s *reservationStorage) List() ([]*models.Reservation, error) {
	// Implementation here
	return nil, nil
}

func (s *reservationStorage) Create(reservation *models.Reservation) (*models.Reservation, error) {
	// Implementation here
	return nil, nil
}

func (s *reservationStorage) Update(reservation *models.Reservation) (*models.Reservation, error) {
	// Implementation here
	return nil, nil
}

func (s *reservationStorage) Delete(id uuid.UUID) error {
	// Implementation here
	return nil
}

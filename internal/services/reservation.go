package service

import (
	"context"
	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
)

type reservationService struct {
	reservationStorage storage.ReservationStorage
}

func NewReservationService(reservationStorage storage.ReservationStorage) ReservationService {
	return &reservationService{reservationStorage: reservationStorage}
}

func (s *reservationService) Get(ctx context.Context, id string) (*models.Reservation, error) {
	return nil, nil
}

func (s *reservationService) List(ctx context.Context, filter storage.ListReservationFilter) ([]*models.Reservation, error) {
	return nil, nil
}

func (s *reservationService) Create(ctx context.Context, reservation *dto.CreateReservationRequest) (*models.Reservation, error) {
	return nil, nil
}

func (s *reservationService) Update(ctx context.Context, reservation *models.Reservation) (*models.Reservation, error) {
	return nil, nil
}

func (s *reservationService) Delete(ctx context.Context, id string) error {
	return nil
}

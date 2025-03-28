package service

import (
	"go-booking/internal/storage"
)

type reservationService struct {
	reservationStorage storage.ReservationStorage
}

func NewReservationService(reservationStorage storage.ReservationStorage) ReservationService {
	return &reservationService{reservationStorage: reservationStorage}
}

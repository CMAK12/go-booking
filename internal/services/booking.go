package service

import (
	"context"
	"log"

	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
)

type bookingService struct {
	bookingStorage storage.BookingStorage
}

func NewBookingService(bookingStorage storage.BookingStorage) BookingService {

	return &bookingService{
		bookingStorage: bookingStorage,
	}
}

func (s *bookingService) List(ctx context.Context, filter storage.ListBookingFilter) ([]models.Booking, error) {
	bookings, err := s.bookingStorage.List(ctx, filter)
	if err != nil {
		log.Println("failed to list bookings:", err)
		return nil, err
	}

	return bookings, nil
}

func (s *bookingService) Create(ctx context.Context, dto dto.CreateBookingRequest) (models.Booking, error) {
	booking, err := s.bookingStorage.Create(ctx, models.NewBooking(dto))
	if err != nil {
		log.Println("failed to create booking:", err)
		return models.Booking{}, err
	}

	return booking, nil
}

func (s *bookingService) Update(ctx context.Context, id string, dto dto.UpdateBookingRequest) (models.Booking, error) {
	updatedBooking, err := s.bookingStorage.Update(ctx, id, models.NewBookingFromDTO(id, dto))
	if err != nil {
		log.Println("failed to update booking:", err)
		return models.Booking{}, err
	}

	return updatedBooking, nil
}

func (s *bookingService) Delete(ctx context.Context, id string) error {
	err := s.bookingStorage.Delete(ctx, id)
	if err != nil {
		log.Println("failed to delete booking:", err)
		return err
	}
	return nil
}

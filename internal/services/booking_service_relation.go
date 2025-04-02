package service

import (
	"context"
	"log"

	d "go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
)

type bookingServiceRelationService struct {
	storage storage.BookingServiceRelationStorage
}

func NewBookingServiceRelationService(storage storage.BookingServiceRelationStorage) BookingServiceRelationService {
	return &bookingServiceRelationService{
		storage: storage,
	}
}

func (s *bookingServiceRelationService) List(ctx context.Context, filter storage.ListBookingServiceRelationFilter) ([]models.BookingServiceRelation, error) {
	relations, err := s.storage.List(ctx, filter)
	if err != nil {
		log.Println("Error listing booking service relations:", err)
		return nil, err
	}

	return relations, nil
}

func (s *bookingServiceRelationService) Create(ctx context.Context, dto d.CreateBookingServiceRelationRequest) (d.CreateBookingServiceRelationResponse, error) {
	bookingServiceRelation, err := s.storage.Create(ctx, dto.BookingID, dto.ServiceID)
	if err != nil {
		log.Println("Error creating booking service relation:", err)
		return d.CreateBookingServiceRelationResponse{}, err
	}

	return bookingServiceRelation, nil
}

func (s *bookingServiceRelationService) Delete(ctx context.Context, bookingID, serviceID string) error {
	err := s.storage.Delete(ctx, bookingID, serviceID)
	if err != nil {
		log.Println("Error deleting booking service relation:", err)
		return err
	}

	return nil
}

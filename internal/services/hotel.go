package service

import (
	"context"
	"log"

	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
)

type hotelService struct {
	hotelStorage storage.HotelStorage
}

func NewHotelService(hotelStorage storage.HotelStorage) HotelService {
	return &hotelService{hotelStorage: hotelStorage}
}

func (s *hotelService) List(ctx context.Context, filter storage.ListHotelFilter) ([]models.Hotel, int64, error) {
	hotels, count, err := s.hotelStorage.List(ctx, filter)
	if err != nil {
		log.Println("failed to list hotels:", err)
		return nil, 0, err
	}

	return hotels, count, nil
}

func (s *hotelService) Create(ctx context.Context, dto dto.CreateHotelRequest) (models.Hotel, error) {
	hotel := models.NewHotel(
		dto.Name,
		dto.Address,
		dto.City,
		dto.Description,
		float64(dto.Rating),
	)

	newHotel, err := s.hotelStorage.Create(ctx, hotel)
	if err != nil {
		log.Println("failed to create hotel:", err)
		return models.Hotel{}, err
	}

	return newHotel, nil
}

func (s *hotelService) Update(ctx context.Context, id string, hotel models.Hotel) (models.Hotel, error) {
	hotel, err := s.hotelStorage.Update(ctx, id, hotel)
	if err != nil {
		log.Println("failed to update hotel:", err)
		return models.Hotel{}, err
	}

	return hotel, nil
}

func (s *hotelService) Delete(ctx context.Context, id string) error {
	err := s.hotelStorage.Delete(ctx, id)
	if err != nil {
		log.Println("failed to delete hotel:", err)
		return err
	}

	return nil
}

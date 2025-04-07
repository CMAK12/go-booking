package service

import (
	"context"

	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
)

type (
	UserService interface {
		List(ctx context.Context, filter storage.ListUserFilter) ([]models.User, error)
		Create(ctx context.Context, dto dto.CreateUserRequest) (models.User, error)
		Update(ctx context.Context, id string, user models.User) (models.User, error)
		Delete(ctx context.Context, id string) error
	}

	BookingService interface {
		List(ctx context.Context, filter storage.ListBookingFilter) ([]dto.ListBookingResponse, error)
		Create(ctx context.Context, dto dto.CreateBookingRequest) (models.Booking, error)
		Update(ctx context.Context, id string, dto dto.UpdateBookingRequest) (models.Booking, error)
		Delete(ctx context.Context, id string) error
	}

	HotelService interface {
		List(ctx context.Context, filter storage.ListHotelFilter) ([]models.Hotel, error)
		Create(ctx context.Context, dto dto.CreateHotelRequest) (models.Hotel, error)
		Update(ctx context.Context, id string, hotel models.Hotel) (models.Hotel, error)
		Delete(ctx context.Context, id string) error
	}

	RoomService interface {
		List(ctx context.Context, filter storage.ListRoomFilter) ([]dto.ListRoomResponse, error)
		Create(ctx context.Context, dto dto.CreateRoomRequest) (models.Room, error)
		Update(ctx context.Context, id string, room models.Room) (models.Room, error)
		Delete(ctx context.Context, id string) error
	}

	ExtraServiceService interface {
		List(ctx context.Context, filter storage.ListExtraServiceFilter) ([]models.ExtraService, error)
		Create(ctx context.Context, dto dto.CreateExtraServiceRequest) (models.ExtraService, error)
		Update(ctx context.Context, id string, extraService models.ExtraService) (models.ExtraService, error)
		Delete(ctx context.Context, id string) error
	}
)

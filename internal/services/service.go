package service

import (
	"context"

	"go-booking/internal/dto"
	"go-booking/internal/filter"
	"go-booking/internal/models"
)

type (
	UserService interface {
		List(ctx context.Context, filter filter.ListUserFilter) ([]models.User, int64, error)
		Create(ctx context.Context, dto dto.CreateUserRequest) (models.User, error)
		Update(ctx context.Context, id string, dto dto.UpdateUserRequest) (models.User, error)
		Delete(ctx context.Context, id string) error
	}

	BookingService interface {
		List(ctx context.Context, filter filter.ListBookingFilter) ([]dto.ListBookingResponse, int64, error)
		Create(ctx context.Context, dto dto.CreateBookingRequest) (models.Booking, error)
		Update(ctx context.Context, id string, dto dto.UpdateBookingRequest) (models.Booking, error)
		Delete(ctx context.Context, id string) error
	}

	HotelService interface {
		List(ctx context.Context, filter filter.ListHotelFilter) ([]models.Hotel, int64, error)
		Create(ctx context.Context, dto dto.CreateHotelRequest) (models.Hotel, error)
		Update(ctx context.Context, id string, hotel models.Hotel) (models.Hotel, error)
		Delete(ctx context.Context, id string) error
	}

	RoomService interface {
		List(ctx context.Context, filter filter.ListRoomFilter) ([]dto.ListRoomResponse, int64, error)
		Create(ctx context.Context, dto dto.CreateRoomRequest) (models.Room, error)
		Update(ctx context.Context, id string, room models.Room) (models.Room, error)
		Delete(ctx context.Context, id string) error
	}

	ExtraServiceService interface {
		List(ctx context.Context, filter filter.ListExtraServiceFilter) ([]models.ExtraService, int64, error)
		Create(ctx context.Context, dto dto.CreateExtraServiceRequest) (models.ExtraService, error)
		Update(ctx context.Context, id string, extraService models.ExtraService) (models.ExtraService, error)
		Delete(ctx context.Context, id string) error
	}
)

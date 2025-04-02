package service

import (
	"context"

	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
)

type (
	UserService interface {
		Get(ctx context.Context, id string) (models.User, error)
		List(ctx context.Context, filter storage.ListUserFilter) ([]models.User, error)
		Create(ctx context.Context, dto dto.CreateUserRequest) (models.User, error)
		Update(ctx context.Context, id string, user models.User) (models.User, error)
		Delete(ctx context.Context, id string) error
	}

	BookingService interface {
		Get(ctx context.Context, id string) (models.Booking, error)
		List(ctx context.Context, filter storage.ListBookingFilter) ([]models.Booking, error)
		Create(ctx context.Context, dto dto.CreateBookingRequest) (models.Booking, error)
		Update(ctx context.Context, id string, dto dto.UpdateBookingRequest) (models.Booking, error)
		Delete(ctx context.Context, id string) error
	}

	HotelService interface {
		Get(ctx context.Context, id string) (models.Hotel, error)
		List(ctx context.Context, filter storage.ListHotelFilter) ([]models.Hotel, error)
		Create(ctx context.Context, dto dto.CreateHotelRequest) (models.Hotel, error)
		Update(ctx context.Context, id string, hotel models.Hotel) (models.Hotel, error)
		Delete(ctx context.Context, id string) error
	}

	RoomService interface {
		Get(ctx context.Context, id string) (models.Room, error)
		List(ctx context.Context, filter storage.ListRoomFilter) ([]models.Room, error)
		Create(ctx context.Context, dto dto.CreateRoomRequest) (models.Room, error)
		Update(ctx context.Context, id string, room models.Room) (models.Room, error)
		Delete(ctx context.Context, id string) error
	}

	ExtraServiceService interface {
		Get(ctx context.Context, id string) (models.ExtraService, error)
		List(ctx context.Context, filter storage.ListExtraServiceFilter) ([]models.ExtraService, error)
		Create(ctx context.Context, dto dto.CreateExtraServiceRequest) (models.ExtraService, error)
		Update(ctx context.Context, id string, extraService models.ExtraService) (models.ExtraService, error)
		Delete(ctx context.Context, id string) error
	}

	BookingServiceRelationService interface {
		List(ctx context.Context, filter storage.ListBookingServiceRelationFilter) ([]models.BookingServiceRelation, error)
		Create(ctx context.Context, dto dto.CreateBookingServiceRelationRequest) (dto.CreateBookingServiceRelationResponse, error)
		Delete(ctx context.Context, bookingID, serviceID string) error
	}

	DiscountService interface {
		Get(ctx context.Context, id string) (models.Discount, error)
		List(ctx context.Context, filter storage.ListDiscountFilter) ([]models.Discount, error)
		Create(ctx context.Context, dto dto.CreateDiscountRequest) (models.Discount, error)
		Update(ctx context.Context, id string, discount models.Discount) (models.Discount, error)
		Delete(ctx context.Context, id string) error
	}
)

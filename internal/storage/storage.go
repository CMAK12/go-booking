package storage

import (
	"context"

	"go-booking/internal/dto"
	"go-booking/internal/models"
)

const (
	userTable         = "users"
	bookingTable      = "bookings"
	roomTable         = "rooms"
	hotelTable        = "hotels"
	extraServiceTable = "services"
)

type (
	UserStorage interface {
		List(ctx context.Context, filter dto.ListUserFilter) ([]models.User, int64, error)
		Create(ctx context.Context, user models.User) (models.User, error)
		Update(ctx context.Context, id string, user models.User) (models.User, error)
		Delete(ctx context.Context, id string) error
	}

	BookingStorage interface {
		List(ctx context.Context, filter dto.ListBookingFilter) ([]models.Booking, int64, error)
		Create(ctx context.Context, booking models.Booking) (models.Booking, error)
		Update(ctx context.Context, id string, booking models.Booking) (models.Booking, error)
		Delete(ctx context.Context, id string) error
	}

	HotelStorage interface {
		List(ctx context.Context, filter dto.ListHotelFilter) ([]models.Hotel, int64, error)
		Create(ctx context.Context, hotel models.Hotel) (models.Hotel, error)
		Update(ctx context.Context, id string, hotel models.Hotel) (models.Hotel, error)
		Delete(ctx context.Context, id string) error
	}

	RoomStorage interface {
		List(ctx context.Context, filter dto.ListRoomFilter) ([]models.Room, int64, error)
		Create(ctx context.Context, room models.Room) (models.Room, error)
		Update(ctx context.Context, id string, room models.Room) (models.Room, error)
		Delete(ctx context.Context, id string) error
	}

	ExtraServiceStorage interface {
		List(ctx context.Context, filter dto.ListExtraServiceFilter) ([]models.ExtraService, int64, error)
		Create(ctx context.Context, extraService models.ExtraService) (models.ExtraService, error)
		Update(ctx context.Context, id string, extraService models.ExtraService) (models.ExtraService, error)
		Delete(ctx context.Context, id string) error
	}
)

package storage

import (
	"context"

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
		List(ctx context.Context, filter ListUserFilter) ([]models.User, error)
		Create(ctx context.Context, user models.User) (models.User, error)
		Update(ctx context.Context, id string, user models.User) (models.User, error)
		Delete(ctx context.Context, id string) error
	}

	BookingStorage interface {
		List(ctx context.Context, filter ListBookingFilter) ([]models.Booking, error)
		Create(ctx context.Context, booking models.Booking) (models.Booking, error)
		Update(ctx context.Context, id string, booking models.Booking) (models.Booking, error)
		Delete(ctx context.Context, id string) error
	}

	HotelStorage interface {
		List(ctx context.Context, filter ListHotelFilter) ([]models.Hotel, error)
		Create(ctx context.Context, hotel models.Hotel) (models.Hotel, error)
		Update(ctx context.Context, id string, hotel models.Hotel) (models.Hotel, error)
		Delete(ctx context.Context, id string) error
	}

	RoomStorage interface {
		List(ctx context.Context, filter ListRoomFilter) ([]models.Room, error)
		Create(ctx context.Context, room models.Room) (models.Room, error)
		Update(ctx context.Context, id string, room models.Room) (models.Room, error)
		Delete(ctx context.Context, id string) error
	}

	ExtraServiceStorage interface {
		List(ctx context.Context, filter ListExtraServiceFilter) ([]models.ExtraService, error)
		Create(ctx context.Context, extraService models.ExtraService) (models.ExtraService, error)
		Update(ctx context.Context, id string, extraService models.ExtraService) (models.ExtraService, error)
		Delete(ctx context.Context, id string) error
	}
)

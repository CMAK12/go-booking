package storage

import (
	"context"
	"go-booking/internal/dto"
	"go-booking/internal/models"
)

const (
	userTable                   = "users"
	bookingTable                = "bookings"
	roomTable                   = "rooms"
	hotelTable                  = "hotels"
	extraServiceTable           = "services"
	bookingServiceRelationTable = "booking_services"
	discountTable               = "discounts"
)

type (
	UserStorage interface {
		Get(ctx context.Context, id string) (models.User, error)
		List(ctx context.Context, filter ListUserFilter) ([]models.User, error)
		Create(ctx context.Context, user models.User) (models.User, error)
		Update(ctx context.Context, id string, user models.User) (models.User, error)
		Delete(ctx context.Context, id string) error
	}

	BookingStorage interface {
		Get(ctx context.Context, id string) (models.Booking, error)
		List(ctx context.Context, filter ListBookingFilter) ([]models.Booking, error)
		Create(ctx context.Context, booking models.Booking) (models.Booking, error)
		Update(ctx context.Context, id string, booking models.Booking) (models.Booking, error)
		Delete(ctx context.Context, id string) error
	}

	HotelStorage interface {
		Get(ctx context.Context, id string) (models.Hotel, error)
		List(ctx context.Context, filter ListHotelFilter) ([]models.Hotel, error)
		Create(ctx context.Context, hotel models.Hotel) (models.Hotel, error)
		Update(ctx context.Context, id string, hotel models.Hotel) (models.Hotel, error)
		Delete(ctx context.Context, id string) error
	}

	RoomStorage interface {
		Get(ctx context.Context, id string) (models.Room, error)
		List(ctx context.Context, filter ListRoomFilter) ([]models.Room, error)
		Create(ctx context.Context, room models.Room) (models.Room, error)
		Update(ctx context.Context, id string, room models.Room) (models.Room, error)
		Delete(ctx context.Context, id string) error
	}

	ExtraServiceStorage interface {
		Get(ctx context.Context, id string) (models.ExtraService, error)
		List(ctx context.Context, filter ListExtraServiceFilter) ([]models.ExtraService, error)
		Create(ctx context.Context, extraService models.ExtraService) (models.ExtraService, error)
		Update(ctx context.Context, id string, extraService models.ExtraService) (models.ExtraService, error)
		Delete(ctx context.Context, id string) error
	}

	BookingServiceRelationStorage interface {
		List(ctx context.Context, filter ListBookingServiceRelationFilter) ([]models.BookingServiceRelation, error)
		Create(ctx context.Context, bookingID, serviceID string) (dto.CreateBookingServiceRelationResponse, error)
		Delete(ctx context.Context, bookingID, serviceID string) error
	}

	DiscountStorage interface {
		Get(ctx context.Context, id string) (models.Discount, error)
		List(ctx context.Context, filter ListDiscountFilter) ([]models.Discount, error)
		Create(ctx context.Context, discount models.Discount) (models.Discount, error)
		Update(ctx context.Context, id string, discount models.Discount) (models.Discount, error)
		Delete(ctx context.Context, id string) error
	}
)

package models

import (
	"time"

	"go-booking/internal/consts"
	"go-booking/internal/dto"

	"github.com/google/uuid"
)

type Booking struct {
	ID        string    `json:"id"`
	User      User      `json:"user"`
	Room      Room      `json:"room"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Status    string    `json:"status"` // pending, confirmed, cancelled
}

func NewBooking(dto dto.CreateBookingRequest) Booking {
	startDate, _ := time.Parse(time.RFC3339, dto.StartDate)
	endDate, _ := time.Parse(time.RFC3339, dto.EndDate)

	return Booking{
		ID: uuid.NewString(),
		User: User{
			ID: dto.UserID,
		},
		Room: Room{
			ID: dto.RoomID,
			Hotel: Hotel{
				ID: dto.HotelID,
			},
		},
		StartDate: startDate,
		EndDate:   endDate,
		Status:    consts.BookingStatusPending,
	}
}

func NewBookingFromDTO(id string, dto dto.UpdateBookingRequest) Booking {
	startDate, _ := time.Parse(time.RFC3339, dto.StartDate)
	endDate, _ := time.Parse(time.RFC3339, dto.EndDate)

	return Booking{
		ID: id,
		User: User{
			ID: dto.UserID,
		},
		Room: Room{
			ID: dto.RoomID,
		},
		StartDate: startDate,
		EndDate:   endDate,
		Status:    dto.Status,
	}
}

package models

import (
	"time"

	"github.com/google/uuid"
)

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "pending"
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCancelled BookingStatus = "cancelled"
)

type Booking struct {
	ID        string        `json:"id"`
	UserID    string        `json:"user_id"`
	RoomID    string        `json:"room_id"`
	HotelID   string        `json:"hotel_id"`
	StartDate time.Time     `json:"start_date"`
	EndDate   time.Time     `json:"end_date"`
	Status    BookingStatus `json:"status"` // pending, confirmed, cancelled
}

func NewBooking(
	userID string,
	roomID string,
	hotelID string,
	startDateStr string,
	endDateStr string,
) Booking {
	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, startDateStr)
	endDate, _ := time.Parse(layout, endDateStr)

	return Booking{
		ID:        uuid.NewString(),
		UserID:    userID,
		RoomID:    roomID,
		HotelID:   hotelID,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    BookingStatusPending,
	}
}

func NewBookingFromDTO(
	id string,
	userID string,
	roomID string,
	startDateStr string,
	endDateStr string,
	status string,
) Booking {
	startDate, _ := time.Parse(time.RFC3339, startDateStr)
	endDate, _ := time.Parse(time.RFC3339, endDateStr)

	return Booking{
		ID:        id,
		UserID:    userID,
		RoomID:    roomID,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    BookingStatus(status),
	}
}

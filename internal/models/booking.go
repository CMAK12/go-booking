package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "pending"
	BookingStatusConfirmed BookingStatus = "confirmed"
	BookingStatusCancelled BookingStatus = "cancelled"

	timeLayout = "2006-01-02"
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
) (Booking, error) {
	startDate, err := time.Parse(timeLayout, startDateStr)
	if err != nil {
		return Booking{}, fmt.Errorf("error parsing start date: %w", err)
	}
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 10, 0, 0, 0, startDate.Location())

	endDate, err := time.Parse(timeLayout, endDateStr)
	if err != nil {
		return Booking{}, fmt.Errorf("error parsing end date: %w", err)
	}
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 10, 0, 0, 0, endDate.Location())

	return Booking{
		ID:        uuid.NewString(),
		UserID:    userID,
		RoomID:    roomID,
		HotelID:   hotelID,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    BookingStatusPending,
	}, nil
}

func NewBookingFromDTO(
	id string,
	userID string,
	roomID string,
	startDateStr string,
	endDateStr string,
	status string,
) (Booking, error) {
	startDate, err := time.Parse(timeLayout, startDateStr)
	if err != nil {
		return Booking{}, fmt.Errorf("error parsing start date: %w", err)
	}
	startDate = time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 10, 0, 0, 0, startDate.Location())

	endDate, err := time.Parse(timeLayout, endDateStr)
	if err != nil {
		return Booking{}, fmt.Errorf("error parsing end date: %w", err)
	}
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 10, 0, 0, 0, endDate.Location())

	return Booking{
		ID:        id,
		UserID:    userID,
		RoomID:    roomID,
		StartDate: startDate,
		EndDate:   endDate,
		Status:    BookingStatus(status),
	}, nil
}

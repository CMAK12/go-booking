package filter

import "go-booking/internal/models"

type ListBookingFilter struct {
	ID        string               `schema:"id"`
	RoomID    string               `schema:"room_id"`
	UserID    string               `schema:"user_id"`
	StartDate string               `schema:"start_date"`
	EndDate   string               `schema:"end_date"`
	Status    models.BookingStatus `schema:"status"`
}

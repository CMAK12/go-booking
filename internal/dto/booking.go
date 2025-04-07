package dto

import (
	"go-booking/internal/models"
	"time"
)

type CreateBookingRequest struct {
	HotelID   string `json:"hotel_id" validate:"required,uuid"`
	RoomID    string `json:"room_id" validate:"required,uuid"`
	UserID    string `json:"user_id" validate:"required,uuid"`
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
}

type UpdateBookingRequest struct {
	RoomID    string `json:"room_id" validate:"required,uuid"`
	UserID    string `json:"user_id" validate:"required,uuid"`
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date" validate:"required"`
	Status    string `json:"status" validate:"required,oneof=pending confirmed cancelled"`
}

type ListBookingResponse struct {
	ID        string           `json:"id"`
	User      models.User      `json:"user"`
	Room      ListRoomResponse `json:"room"`
	Hotel     models.Hotel     `json:"hotel"`
	StartDate time.Time        `json:"start_date"`
	EndDate   time.Time        `json:"end_date"`
	Status    string           `json:"status"`
}

package dto

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

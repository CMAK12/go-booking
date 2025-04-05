package dto

type CreateRoomRequest struct {
	HotelID  string `json:"hotel_id" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Capacity int    `json:"capacity" validate:"required"`
	Price    int    `json:"price" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}

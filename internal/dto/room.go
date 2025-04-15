package dto

type CreateRoomRequest struct {
	HotelID  string `json:"hotel_id" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Capacity int    `json:"capacity" validate:"required"`
	Price    int    `json:"price" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}

type ListRoomResponse struct {
	ID            string                     `json:"id"`
	HotelID       string                     `json:"hotel_id"`
	ExtraServices []ListExtraServiceResponse `json:"extra_services"`
	Type          string                     `json:"type"`
	Capacity      int                        `json:"capacity"`
	Price         float64                    `json:"price"`
	Quantity      int                        `json:"quantity"`
}

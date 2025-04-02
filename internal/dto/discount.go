package dto

type CreateDiscountRequest struct {
	HotelID string  `json:"hotel_id"`
	Name    string  `json:"name"`
	Amount  float64 `json:"amount"`
	Active  bool    `json:"active"`
}

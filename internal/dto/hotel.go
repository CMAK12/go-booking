package dto

type CreateHotelRequest struct {
	Name        string `json:"name" validate:"required"`
	Address     string `json:"address" validate:"required"`
	City        string `json:"city" validate:"required"`
	Description string `json:"description"`
	Rating      int    `json:"rating" validate:"required"`
}

type HotelResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Rating      int    `json:"rating"`
}

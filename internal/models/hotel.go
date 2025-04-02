package models

import (
	"go-booking/internal/dto"

	"github.com/google/uuid"
)

type Hotel struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	City        string  `json:"city"`
	Description string  `json:"description"`
	Rating      float64 `json:"rating"`
}

func NewHotel(dto dto.CreateHotelRequest) Hotel {
	return Hotel{
		ID:          uuid.NewString(),
		Name:        dto.Name,
		Address:     dto.Address,
		City:        dto.City,
		Description: dto.Description,
		Rating:      float64(dto.Rating),
	}
}

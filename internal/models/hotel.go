package models

import (
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

func NewHotel(
	name string,
	address string,
	city string,
	description string,
	rating float64,
) Hotel {
	return Hotel{
		ID:          uuid.NewString(),
		Name:        name,
		Address:     address,
		City:        city,
		Description: description,
		Rating:      rating,
	}
}

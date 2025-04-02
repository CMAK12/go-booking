package models

import (
	"go-booking/internal/dto"

	"github.com/google/uuid"
)

type Discount struct {
	ID     string  `json:"id"`
	Hotel  Hotel   `json:"room"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Active bool    `json:"active"`
}

func NewDiscount(dto dto.CreateDiscountRequest) Discount {
	return Discount{
		ID:     uuid.NewString(),
		Hotel:  Hotel{ID: dto.HotelID},
		Name:   dto.Name,
		Amount: dto.Amount,
		Active: dto.Active,
	}
}

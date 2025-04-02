package models

import (
	"go-booking/internal/dto"

	"github.com/google/uuid"
)

type Room struct {
	ID        string  `json:"id"`
	Hotel     Hotel   `json:"hotel"`
	Type      string  `json:"type"`
	Capacity  int     `json:"capacity"`
	Price     float64 `json:"price"`
	Available bool    `json:"available"`
}

func NewRoom(dto dto.CreateRoomRequest) Room {
	return Room{
		ID:        uuid.NewString(),
		Hotel:     Hotel{ID: dto.HotelID},
		Type:      dto.Type,
		Capacity:  dto.Capacity,
		Price:     float64(dto.Price),
		Available: dto.Available,
	}
}

package models

import (
	"github.com/google/uuid"
)

type Room struct {
	ID              string   `json:"id"`
	HotelID         string   `json:"hotel_id"`
	ExtraServiceIDs []string `json:"extra_service_ids"`
	Type            string   `json:"type"`
	Capacity        int      `json:"capacity"`
	Price           float64  `json:"price"`
	Quantity        int      `json:"quantity"`
}

func NewRoom(
	hotelID string,
	typeRoom string,
	capacity int,
	price float64,
	quantity int,
) Room {
	return Room{
		ID:       uuid.NewString(),
		HotelID:  hotelID,
		Type:     typeRoom,
		Capacity: capacity,
		Price:    price,
		Quantity: quantity,
	}
}

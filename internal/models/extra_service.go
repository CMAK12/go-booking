package models

import (
	"github.com/google/uuid"
)

type ExtraService struct {
	ID     string `json:"id"`
	RoomID string `json:"room_id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}

func NewExtraService(
	roomID string,
	name string,
	price int,
) ExtraService {
	return ExtraService{
		ID:     uuid.NewString(),
		RoomID: roomID,
		Name:   name,
		Price:  price,
	}
}

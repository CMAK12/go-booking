package models

import (
	"go-booking/internal/dto"

	"github.com/google/uuid"
)

type ExtraService struct {
	ID     string `json:"id"`
	RoomID string `json:"-"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}

func NewExtraService(dto dto.CreateExtraServiceRequest) ExtraService {
	return ExtraService{
		ID:     uuid.NewString(),
		RoomID: dto.RoomID,
		Name:   dto.Name,
		Price:  dto.Price,
	}
}

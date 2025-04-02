package models

import (
	"go-booking/internal/dto"

	"github.com/google/uuid"
)

type ExtraService struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func NewExtraService(dto dto.CreateExtraServiceRequest) ExtraService {
	return ExtraService{
		ID:    uuid.NewString(),
		Name:  dto.Name,
		Price: dto.Price,
	}
}

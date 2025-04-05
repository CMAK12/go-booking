package dto

type CreateExtraServiceRequest struct {
	RoomID string `json:"room_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Price  int    `json:"price" validate:"required"`
}

package dto

type CreateExtraServiceRequest struct {
	RoomID string `json:"room_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Price  int    `json:"price" validate:"required"`
}

type ListExtraServiceResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

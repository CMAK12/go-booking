package dto

type CreateReservationRequest struct {
	FirstName     string `json:"first_name" validate:"required,min=3,max=20"`
	LastName      string `json:"last_name" validate:"required,min=3,max=20"`
	Date          string `json:"date" validate:"required"`
	GuestQuantity int    `json:"guest_quantity" validate:"required,min=1,max=10"`
	City          string `json:"city" validate:"required,min=3,max=20"`
}

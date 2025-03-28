package models

type Reservation struct {
	ID            int    `json:"id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Date          string `json:"date"`
	GuestQuantity int    `json:"guest_quantity"`
	City          string `json:"city"`
}

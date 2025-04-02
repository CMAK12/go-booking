package models

type BookingServiceRelation struct {
	Booking Booking      `json:"booking"`
	Extra   ExtraService `json:"extra_service"`
}

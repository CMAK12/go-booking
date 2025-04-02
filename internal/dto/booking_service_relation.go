package dto

type (
	CreateBookingServiceRelationRequest struct {
		BookingID string `json:"booking_id"`
		ServiceID string `json:"service_id"`
	}

	CreateBookingServiceRelationResponse struct {
		BookingID string `json:"booking_id"`
		ServiceID string `json:"service_id"`
	}
)

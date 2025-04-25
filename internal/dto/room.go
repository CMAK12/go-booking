package dto

type CreateRoomRequest struct {
	HotelID  string `json:"hotel_id" validate:"required"`
	Type     string `json:"type" validate:"required"`
	Capacity int    `json:"capacity" validate:"required"`
	Price    int    `json:"price" validate:"required"`
	Quantity int    `json:"quantity" validate:"required"`
}

type ListRoomResponse struct {
	ID            string                     `json:"id"`
	HotelID       string                     `json:"hotel_id"`
	ExtraServices []ListExtraServiceResponse `json:"extra_services"`
	Type          string                     `json:"type"`
	Capacity      int                        `json:"capacity"`
	Price         float64                    `json:"price"`
	Quantity      int                        `json:"quantity"`
}

type ListRoomFilter struct {
	ID          string   `schema:"id"`
	IDs         []string `schema:"ids,omitempty"`
	HotelID     string   `schema:"hotel_id"`
	Name        string   `schema:"name"`
	Description string   `schema:"description"`
	Price       int      `schema:"price"`
	Capacity    int      `schema:"capacity"`
	Quantity    int      `schema:"quantity"`
	ExcludeIDs  []string `schema:"exclude_ids"`
	Take        int64    `schema:"take"`
	Skip        int64    `schema:"skip"`
}

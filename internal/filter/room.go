package filter

type ListRoomFilter struct {
	ID          string   `schema:"id"`
	IDs         []string `schema:"ids,omitempty"`
	HotelID     string   `schema:"hotel_id"`
	Name        string   `schema:"name"`
	Description string   `schema:"description"`
	Price       int      `schema:"price"`
	Capacity    int      `schema:"capacity"`
	Quantity    int      `schema:"quantity"`
}

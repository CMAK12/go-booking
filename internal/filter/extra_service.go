package filter

type ListExtraServiceFilter struct {
	ID     string `schema:"id"`
	RoomID string `schema:"room_id"`
	Name   string `schema:"name"`
	Price  int    `schema:"price"`
}

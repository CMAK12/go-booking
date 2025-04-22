package filter

type ListHotelFilter struct {
	ID          string   `schema:"id"`
	IDs         []string `schema:"ids,omitempty"`
	Name        string   `schema:"name"`
	City        string   `schema:"city"`
	Address     string   `schema:"address"`
	Description string   `schema:"description"`
	Rating      float64  `schema:"rating"`
}

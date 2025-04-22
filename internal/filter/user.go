package filter

import "go-booking/internal/models"

type ListUserFilter struct {
	ID       string          `schema:"id"`
	IDs      []string        `schema:"ids,omitempty"`
	Username string          `schema:"username"`
	Email    string          `schema:"email"`
	Role     models.UserRole `schema:"role"`
}

package dto

import "go-booking/internal/models"

type CreateUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type ListUserFilter struct {
	ID       string          `schema:"id"`
	IDs      []string        `schema:"ids,omitempty"`
	Username string          `schema:"username"`
	Email    string          `schema:"email"`
	Role     models.UserRole `schema:"role"`
}

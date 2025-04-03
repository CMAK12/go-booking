package models

import (
	"time"

	"go-booking/internal/dto"

	"github.com/google/uuid"
)

type UserRole string

const (
	RoleGuest   UserRole = "guest"
	RoleAdmin   UserRole = "admin"
	RoleManager UserRole = "manager"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(dto dto.CreateUserRequest) User {
	return User{
		ID:        uuid.NewString(),
		Email:     dto.Email,
		Name:      dto.Name,
		Password:  dto.Password,
		CreatedAt: time.Now(),
		Role:      RoleGuest,
	}
}

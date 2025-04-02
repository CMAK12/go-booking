package models

import (
	"go-booking/internal/consts"
	"go-booking/internal/dto"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"` // guest, admin, manager
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(dto dto.CreateUserRequest) User {
	return User{
		ID:        uuid.NewString(),
		Email:     dto.Email,
		Name:      dto.Name,
		Password:  dto.Password,
		CreatedAt: time.Now(),
		Role:      consts.RoleGuest,
	}
}

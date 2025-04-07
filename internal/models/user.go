package models

import (
	"time"

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

func NewUser(
	email string,
	name string,
	password string,
) User {
	return User{
		ID:        uuid.NewString(),
		Email:     email,
		Name:      name,
		Password:  password,
		CreatedAt: time.Now(),
		Role:      RoleGuest,
	}
}

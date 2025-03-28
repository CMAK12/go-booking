package storage

import (
	"database/sql"

	"go-booking/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type userStorage struct {
	db 			*sql.DB
	builder sq.StatementBuilderType
}

func NewUserStorage(db *sql.DB) UserStorage {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &userStorage{
		db: db,
		builder: builder,
	}
}

func (s *userStorage) Get(id uuid.UUID) (*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userStorage) List() ([]*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userStorage) Create(user *models.User) (*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userStorage) Update(user *models.User) (*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userStorage) Delete(id uuid.UUID) error {
	// Implementation here
	return nil
}

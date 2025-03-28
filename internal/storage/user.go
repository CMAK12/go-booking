package storage

import (
	"context"

	"go-booking/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	userStorage struct {
		db 			*pgxpool.Pool
		builder sq.StatementBuilderType
	}

	ListUserFilter struct {
		ID string `json:"id"`
		Username string `json:"username"`
		Email string `json:"email"`
		Role string `json:"role"`
	}
)

func NewUserStorage(db *pgxpool.Pool) UserStorage {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &userStorage{
		db: db,
		builder: builder,
	}
}

func (s *userStorage) Get(ctx context.Context, id string) (*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userStorage) List(ctx context.Context, filter ListUserFilter) ([]*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userStorage) Create(ctx context.Context, user *models.User) (*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userStorage) Update(ctx context.Context, user *models.User) (*models.User, error) {
	// Implementation here
	return nil, nil
}

func (s *userStorage) Delete(ctx context.Context, id string) error {
	// Implementation here
	return nil
}

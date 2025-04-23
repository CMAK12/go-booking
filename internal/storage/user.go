package storage

import (
	"context"
	"fmt"

	"go-booking/internal/dto"
	"go-booking/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userStorage struct {
	db      *pgxpool.Pool
	builder sq.StatementBuilderType
}

func NewUserStorage(db *pgxpool.Pool) UserStorage {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &userStorage{
		db:      db,
		builder: builder,
	}
}

func (s *userStorage) List(ctx context.Context, filter dto.ListUserFilter) ([]models.User, int64, error) {
	qb := s.builder.
		Select(
			"id", "name", "email", "role", "created_at",

			"COUNT(*) OVER() AS total_count",
		).
		From(userTable)

	qb = buildSearchUserQuery(qb, filter)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var users []models.User
	var totalCount int64
	for rows.Next() {
		var user models.User
		var count int64
		if err := rows.Scan(
			&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt,
			&count,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}
		totalCount = count
		users = append(users, user)
	}
	return users, totalCount, nil
}

func (s *userStorage) Create(ctx context.Context, user models.User) (models.User, error) {
	query, args, err := s.builder.
		Insert(userTable).
		Columns("id", "name", "email", "password", "role", "created_at").
		Values(user.ID, user.Name, user.Email, user.Password, user.Role, user.CreatedAt).
		ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to execute query: %w", err)
	}
	return user, nil
}

func (s *userStorage) Update(ctx context.Context, id string, user models.User) (models.User, error) {
	if id == "" {
		return models.User{}, fmt.Errorf("user ID is empty")
	}

	query, args, err := s.builder.
		Update(userTable).
		Set("name", user.Name).
		Set("email", user.Email).
		Set("password", user.Password).
		Set("role", user.Role).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING id, created_at").
		ToSql()
	if err != nil {
		return models.User{}, fmt.Errorf("failed to build query: %w", err)
	}

	err = s.db.QueryRow(ctx, query, args...).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to scan row: %w", err)
	}
	return user, nil
}

func (s *userStorage) Delete(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("user ID is empty")
	}

	query, args, err := s.builder.
		Delete(userTable).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}
	return nil
}

func buildSearchUserQuery(qb sq.SelectBuilder, filter dto.ListUserFilter) sq.SelectBuilder {
	if len(filter.IDs) > 0 {
		qb = qb.Where(sq.Eq{"id": filter.IDs})
	}
	if filter.ID != "" {
		qb = qb.Where(sq.Eq{"id": filter.ID})
	}
	if filter.Username != "" {
		qb = qb.Where(sq.Eq{"name": filter.Username})
	}
	if filter.Email != "" {
		qb = qb.Where(sq.Eq{"email": filter.Email})
	}
	if filter.Role != "" {
		qb = qb.Where(sq.Eq{"role": filter.Role})
	}

	return qb
}

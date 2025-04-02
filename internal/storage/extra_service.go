package storage

import (
	"context"
	"fmt"
	"go-booking/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type extraServiceStorage struct {
	db      *pgxpool.Pool
	builder sq.StatementBuilderType
}

type ListExtraServiceFilter struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func NewExtraServiceStorage(db *pgxpool.Pool) ExtraServiceStorage {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &extraServiceStorage{
		db:      db,
		builder: builder,
	}
}

func (s *extraServiceStorage) Get(ctx context.Context, id string) (models.ExtraService, error) {
	if id == "" {
		return models.ExtraService{}, fmt.Errorf("extra service id is empty")
	}

	query, args, err := s.builder.
		Select("id", "name", "price").
		From(extraServiceTable).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return models.ExtraService{}, fmt.Errorf("failed to build query: %w", err)
	}

	var extraService models.ExtraService
	err = s.db.QueryRow(ctx, query, args...).Scan(&extraService.ID, &extraService.Name, &extraService.Price)
	if err != nil {
		return models.ExtraService{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return extraService, nil
}

func (s *extraServiceStorage) List(ctx context.Context, filter ListExtraServiceFilter) ([]models.ExtraService, error) {
	qb := s.builder.
		Select("id", "name", "price").
		From(extraServiceTable)

	if filter.Name != "" {
		qb = qb.Where(sq.Like{"name": "%" + filter.Name + "%"})
	}
	if filter.Price != 0 {
		qb = qb.Where(sq.Eq{"price": filter.Price})
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var extraServices []models.ExtraService
	for rows.Next() {
		var extraService models.ExtraService
		err = rows.Scan(&extraService.ID, &extraService.Name, &extraService.Price)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		extraServices = append(extraServices, extraService)
	}

	return extraServices, nil
}

func (s *extraServiceStorage) Create(ctx context.Context, extraService models.ExtraService) (models.ExtraService, error) {
	query, args, err := s.builder.
		Insert(extraServiceTable).
		Columns("id", "name", "price").
		Values(extraService.ID, extraService.Name, extraService.Price).
		ToSql()
	if err != nil {
		return models.ExtraService{}, fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return models.ExtraService{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return extraService, nil
}

func (s *extraServiceStorage) Update(ctx context.Context, id string, extraService models.ExtraService) (models.ExtraService, error) {
	if id == "" {
		return models.ExtraService{}, fmt.Errorf("extra service id is empty")
	}

	query, args, err := s.builder.
		Update(extraServiceTable).
		Set("name", extraService.Name).
		Set("price", extraService.Price).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return models.ExtraService{}, fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return models.ExtraService{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return extraService, nil
}

func (s *extraServiceStorage) Delete(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("extra service id is empty")
	}

	query, args, err := s.builder.
		Delete(extraServiceTable).
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

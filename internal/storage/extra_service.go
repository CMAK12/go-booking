package storage

import (
	"context"
	"fmt"
	"go-booking/internal/dto"
	"go-booking/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type extraServiceStorage struct {
	db      *pgxpool.Pool
	builder sq.StatementBuilderType
}

func NewExtraServiceStorage(db *pgxpool.Pool) ExtraServiceStorage {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &extraServiceStorage{
		db:      db,
		builder: builder,
	}
}

func (s *extraServiceStorage) List(ctx context.Context, filter dto.ListExtraServiceFilter) ([]models.ExtraService, int64, error) {
	qb := s.builder.
		Select(
			"id", "room_id", "name", "price",

			"COUNT(*) OVER() AS total_count",
		).
		From(extraServiceTable)

	qb = buildSearchExtraServiceQuery(qb, filter)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var extraServices []models.ExtraService
	var totalCount int64
	for rows.Next() {
		var extraService models.ExtraService
		var count int64
		if err = rows.Scan(
			&extraService.ID, &extraService.RoomID, &extraService.Name, &extraService.Price,
			&count,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}

		totalCount = count
		extraServices = append(extraServices, extraService)
	}

	return extraServices, totalCount, nil
}

func (s *extraServiceStorage) Create(ctx context.Context, extraService models.ExtraService) (models.ExtraService, error) {
	query, args, err := s.builder.
		Insert(extraServiceTable).
		Columns("id", "room_id", "name", "price").
		Values(extraService.ID, extraService.RoomID, extraService.Name, extraService.Price).
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
		Set("room_id", extraService.RoomID).
		Set("name", extraService.Name).
		Set("price", extraService.Price).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return models.ExtraService{}, fmt.Errorf("failed to build query: %w", err)
	}

	extraService.ID = id

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

func buildSearchExtraServiceQuery(qb sq.SelectBuilder, filter dto.ListExtraServiceFilter) sq.SelectBuilder {
	if filter.ID != "" {
		qb = qb.Where(sq.Eq{"id": filter.ID})
	}
	if filter.RoomID != "" {
		qb = qb.Where(sq.Eq{"room_id": filter.RoomID})
	}
	if filter.Name != "" {
		qb = qb.Where(sq.Like{"name": "%" + filter.Name + "%"})
	}
	if filter.Price != 0 {
		qb = qb.Where(sq.Eq{"price": filter.Price})
	}

	return qb
}

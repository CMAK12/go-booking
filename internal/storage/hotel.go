package storage

import (
	"context"
	"fmt"

	"go-booking/internal/dto"
	"go-booking/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type hotelStorage struct {
	db      *pgxpool.Pool
	builder sq.StatementBuilderType
}

func NewHotelStorage(db *pgxpool.Pool) HotelStorage {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &hotelStorage{
		db:      db,
		builder: builder,
	}
}

func (s *hotelStorage) List(ctx context.Context, filter dto.ListHotelFilter) ([]models.Hotel, int64, error) {
	qb := s.builder.
		Select(
			"id", "name", "city", "address", "rating", "description",

			"COUNT(*) OVER() AS total_count",
		).
		From(hotelTable)

	qb = buildSearchHotelQuery(qb, filter)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var hotel []models.Hotel
	var totalCount int64
	for rows.Next() {
		var h models.Hotel
		var count int64
		if err := rows.Scan(
			&h.ID, &h.Name, &h.City, &h.Address, &h.Rating, &h.Description,
			&count,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan hotel: %w", err)
		}
		totalCount = count
		hotel = append(hotel, h)
	}
	return hotel, totalCount, nil
}

func (s *hotelStorage) Create(ctx context.Context, hotel models.Hotel) (models.Hotel, error) {
	query, args, err := s.builder.
		Insert(hotelTable).
		Columns("id", "name", "city", "address", "rating", "description").
		Values(hotel.ID, hotel.Name, hotel.City, hotel.Address, hotel.Rating, hotel.Description).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return models.Hotel{}, fmt.Errorf("failed to build query: %w", err)
	}

	err = s.db.QueryRow(ctx, query, args...).Scan(&hotel.ID)
	if err != nil {
		return models.Hotel{}, fmt.Errorf("failed to execute query: %w", err)
	}
	return hotel, nil
}

func (s *hotelStorage) Update(ctx context.Context, id string, hotel models.Hotel) (models.Hotel, error) {
	if id == "" {
		return models.Hotel{}, fmt.Errorf("hotel id is empty")
	}

	query, args, err := s.builder.
		Update(hotelTable).
		Set("name", hotel.Name).
		Set("city", hotel.City).
		Set("address", hotel.Address).
		Set("rating", hotel.Rating).
		Set("description", hotel.Description).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return models.Hotel{}, fmt.Errorf("failed to build query: %w", err)
	}

	err = s.db.QueryRow(ctx, query, args...).Scan(&hotel.ID)
	if err != nil {
		return models.Hotel{}, fmt.Errorf("failed to execute query: %w", err)
	}
	return hotel, nil
}

func (s *hotelStorage) Delete(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("hotel id is empty")
	}

	query, args, err := s.builder.
		Delete(hotelTable).
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

func buildSearchHotelQuery(qb sq.SelectBuilder, filter dto.ListHotelFilter) sq.SelectBuilder {
	if len(filter.IDs) > 0 {
		qb = qb.Where(sq.Eq{"id": filter.IDs})
	}
	if filter.ID != "" {
		qb = qb.Where(sq.Eq{"id": filter.ID})
	}
	if filter.Name != "" {
		qb = qb.Where(sq.Like{"name": "%" + filter.Name + "%"})
	}
	if filter.City != "" {
		qb = qb.Where(sq.Like{"city": "%" + filter.City + "%"})
	}
	if filter.Address != "" {
		qb = qb.Where(sq.Like{"address": "%" + filter.Address + "%"})
	}
	if filter.Description != "" {
		qb = qb.Where(sq.Like{"description": "%" + filter.Description + "%"})
	}
	if filter.Rating != 0 {
		qb = qb.Where(sq.Eq{"rating": filter.Rating})
	}

	return qb
}

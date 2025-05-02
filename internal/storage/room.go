package storage

import (
	"context"
	"fmt"

	"go-booking/internal/consts"
	"go-booking/internal/dto"
	"go-booking/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type roomStorage struct {
	db      *pgxpool.Pool
	builder sq.StatementBuilderType
}

func NewRoomStorage(db *pgxpool.Pool) RoomStorage {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &roomStorage{
		db:      db,
		builder: builder,
	}
}

func (s *roomStorage) List(ctx context.Context, filter dto.ListRoomFilter) ([]models.Room, int64, error) {
	qb := s.builder.
		Select(
			"r.id", "r.hotel_id", "r.type", "r.capacity", "r.price", "r.quantity",
		).
		From(fmt.Sprintf("%s AS r", roomTable))

	qb = buildSearchRoomFilter(qb, filter)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var rooms []models.Room
	var totalCount int64
	for rows.Next() {
		var room models.Room
		if err := rows.Scan(
			&room.ID, &room.HotelID, &room.Type, &room.Capacity, &room.Price, &room.Quantity,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan row: %w", err)
		}
		totalCount++
		rooms = append(rooms, room)
	}

	return rooms, totalCount, nil
}

func (s *roomStorage) Create(ctx context.Context, room models.Room) (models.Room, error) {
	query, args, err := s.builder.
		Insert(roomTable).
		Columns("id", "hotel_id", "type", "capacity", "price", "quantity").
		Values(room.ID, room.HotelID, room.Type, room.Capacity, room.Price, room.Quantity).
		ToSql()
	if err != nil {
		return models.Room{}, fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return models.Room{}, fmt.Errorf("failed to execute query: %w", err)
	}
	return room, nil
}

func (s *roomStorage) Update(ctx context.Context, id string, room models.Room) (models.Room, error) {
	if id == "" {
		return models.Room{}, fmt.Errorf("room id is empty")
	}

	query, args, err := s.builder.
		Update(roomTable).
		Set("type", room.Type).
		Set("capacity", room.Capacity).
		Set("price", room.Price).
		Set("quantity", room.Quantity).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return models.Room{}, fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return models.Room{}, fmt.Errorf("failed to execute query: %w", err)
	}
	return room, nil
}

func (s *roomStorage) Delete(ctx context.Context, id string) error {
	if id == "" {
		return fmt.Errorf("room id is empty")
	}

	query, args, err := s.builder.
		Delete(roomTable).
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

func buildSearchRoomFilter(qb sq.SelectBuilder, filter dto.ListRoomFilter) sq.SelectBuilder {
	if filter.PageSize <= 0 {
		filter.PageSize = consts.DefaultPageSize
	}
	if filter.PageNumber <= 0 {
		filter.PageNumber = consts.DefaultPageNumber
	}

	if len(filter.IDs) > 0 && filter.PageSize > 0 && filter.PageNumber > 1 {
		return qb
	}
	if len(filter.IDs) > 0 {
		qb = qb.Where(sq.Eq{"r.id": filter.IDs})
	}
	if filter.ID != "" {
		qb = qb.Where(sq.Eq{"r.id": filter.ID})
	}
	if filter.HotelID != "" {
		qb = qb.Where(sq.Eq{"r.hotel_id": filter.HotelID})
	}
	if filter.Name != "" {
		qb = qb.Where(sq.Like{"r.name": "%" + filter.Name + "%"})
	}
	if filter.Description != "" {
		qb = qb.Where(sq.Like{"r.description": "%" + filter.Description + "%"})
	}
	if filter.Price > 0 {
		qb = qb.Where(sq.Eq{"r.price": filter.Price})
	}
	if filter.Capacity > 0 {
		qb = qb.Where(sq.Eq{"r.capacity": filter.Capacity})
	}
	if filter.Quantity > 0 {
		qb = qb.Where(sq.Eq{"r.quantity": filter.Quantity})
	}
	if len(filter.ExcludeIDs) > 0 {
		qb = qb.Where(sq.NotEq{"r.id": filter.ExcludeIDs})
	}
	if filter.IDs == nil {
		if filter.PageSize > 0 {
			qb = qb.Limit(uint64(filter.PageSize))
		}
		if filter.PageNumber > 0 {
			qb = qb.Offset(uint64((filter.PageNumber - 1) * filter.PageSize))
		}
	}

	return qb
}

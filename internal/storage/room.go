package storage

import (
	"context"
	"fmt"
	"go-booking/internal/models"
	"strconv"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type roomStorage struct {
	db      *pgxpool.Pool
	builder sq.StatementBuilderType
}

type ListRoomFilter struct {
	ID          string `json:"id"`
	HotelID     string `json:"hotel_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Capacity    int    `json:"capacity"`
	Available   string `json:"available"`
}

func NewRoomStorage(db *pgxpool.Pool) RoomStorage {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &roomStorage{
		db:      db,
		builder: builder,
	}
}

func (s *roomStorage) List(ctx context.Context, filter ListRoomFilter) ([]models.Room, error) {
	qb := s.builder.
		Select("r.id", "r.type", "r.capacity", "r.price", "r.available",
			"h.id", "h.name", "h.address", "h.city", "h.description", "h.rating").
		From(fmt.Sprintf("%s AS r", roomTable)).
		Join(fmt.Sprintf("%s AS h ON r.hotel_id = h.id", hotelTable))

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
	if filter.Available != "" {
		available, err := strconv.ParseBool(filter.Available)
		if err != nil {
			return nil, fmt.Errorf("invalid available value: %w", err)
		}
		qb = qb.Where(sq.Eq{"r.available": available})
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

	var rooms []models.Room
	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID, &room.Type, &room.Capacity, &room.Price, &room.Available,
			&room.Hotel.ID, &room.Hotel.Name, &room.Hotel.Address, &room.Hotel.City, &room.Hotel.Description, &room.Hotel.Rating,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (s *roomStorage) Create(ctx context.Context, room models.Room) (models.Room, error) {
	query, args, err := s.builder.
		Insert(roomTable).
		Columns("id", "hotel_id", "type", "capacity", "price", "available").
		Values(room.ID, room.Hotel.ID, room.Type, room.Capacity, room.Price, room.Available).
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
		Set("available", room.Available).
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

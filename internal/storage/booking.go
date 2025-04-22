package storage

import (
	"context"
	"fmt"

	"go-booking/internal/filter"
	"go-booking/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type bookingStorage struct {
	db      *pgxpool.Pool
	builder sq.StatementBuilderType
}

func NewBookingStorage(db *pgxpool.Pool) BookingStorage {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &bookingStorage{
		db:      db,
		builder: builder,
	}
}

func (s *bookingStorage) List(ctx context.Context, filter filter.ListBookingFilter) ([]models.Booking, int64, error) {
	qb := s.builder.
		Select(
			"bt.id", "bt.user_id", "bt.room_id", "bt.start_date", "bt.end_date", "bt.status",

			"COUNT(*) OVER() AS total_count",
		).
		From(fmt.Sprintf("%s AS bt", bookingTable))

	qb, err := buildSearchBookingQuery(qb, filter)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build search query: %w", err)
	}

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, 0, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var bookings []models.Booking
	var totalCount int64
	for rows.Next() {
		var booking models.Booking
		var count int64
		if err := rows.Scan(
			&booking.ID, &booking.UserID, &booking.RoomID, &booking.StartDate, &booking.EndDate, &booking.Status,
			&count,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to scan booking: %w", err)
		}
		totalCount = count
		bookings = append(bookings, booking)
	}

	return bookings, totalCount, nil
}

func (s *bookingStorage) Create(ctx context.Context, booking models.Booking) (models.Booking, error) {
	query, args, err := s.builder.
		Insert(bookingTable).
		Columns("id", "user_id", "room_id", "start_date", "end_date", "status").
		Values(booking.ID, booking.UserID, booking.RoomID, booking.StartDate, booking.EndDate, booking.Status).
		ToSql()
	if err != nil {
		return models.Booking{}, fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return models.Booking{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return booking, nil
}

func (s *bookingStorage) Update(ctx context.Context, id string, booking models.Booking) (models.Booking, error) {
	if id == "" {
		return models.Booking{}, fmt.Errorf("booking id is empty")
	}

	query, args, err := s.builder.
		Update(bookingTable).
		Set("user_id", booking.UserID).
		Set("room_id", booking.RoomID).
		Set("start_date", booking.StartDate).
		Set("end_date", booking.EndDate).
		Set("status", booking.Status).
		Where(sq.Eq{"id": id}).
		ToSql()

	if err != nil {
		return models.Booking{}, fmt.Errorf("failed to build query: %w", err)
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return models.Booking{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return booking, nil
}

func (s *bookingStorage) Delete(ctx context.Context, id string) error {
	query, args, err := s.builder.
		Delete(bookingTable).
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

func buildSearchBookingQuery(qb sq.SelectBuilder, filter filter.ListBookingFilter) (sq.SelectBuilder, error) {
	if filter.ID != "" {
		qb = qb.Where(sq.Eq{"bt.id": filter.ID})
	}
	if filter.RoomID != "" {
		qb = qb.Where(sq.Eq{"bt.room_id": filter.RoomID})
	}
	if filter.UserID != "" {
		qb = qb.Where(sq.Eq{"bt.user_id": filter.UserID})
	}
	if filter.StartDate != "" && filter.EndDate != "" {
		if filter.StartDate >= filter.EndDate {
			return qb, fmt.Errorf("start date is bigger than end date")
		}
		qb = qb.Where(sq.And{
			sq.Lt{"bt.start_date": filter.EndDate},
			sq.Gt{"bt.end_date": filter.StartDate},
		})
	}
	if filter.Status != "" {
		qb = qb.Where(sq.Eq{"bt.status": filter.Status})
	}

	return qb, nil
}

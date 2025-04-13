package storage

import (
	"context"
	"fmt"

	"go-booking/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type bookingStorage struct {
	db      *pgxpool.Pool
	builder sq.StatementBuilderType
}

type ListBookingFilter struct {
	ID        string               `json:"id"`
	RoomID    string               `json:"room_id"`
	UserID    string               `json:"user_id"`
	StartDate string               `json:"start_date"`
	EndDate   string               `json:"end_date"`
	Status    models.BookingStatus `json:"status"`
}

func NewBookingStorage(db *pgxpool.Pool) BookingStorage {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &bookingStorage{
		db:      db,
		builder: builder,
	}
}

func (s *bookingStorage) List(ctx context.Context, filter ListBookingFilter) ([]models.Booking, error) {
	qb := s.builder.
		Select(
			"bt.id", "bt.user_id", "bt.room_id", "bt.start_date", "bt.end_date", "bt.status",
		).
		From(fmt.Sprintf("%s AS bt", bookingTable))

	qb = buildSearchBookingQuery(qb, filter)

	query, args, err := qb.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build query: %w", err)
	}

	rows, err := s.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		err := rows.Scan(
			&booking.ID, &booking.UserID, &booking.RoomID, &booking.StartDate, &booking.EndDate, &booking.Status,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan booking: %w", err)
		}
		bookings = append(bookings, booking)
	}
	return bookings, nil
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

func buildSearchBookingQuery(qb sq.SelectBuilder, filter ListBookingFilter) sq.SelectBuilder {
	if filter.ID != "" {
		qb = qb.Where(sq.Eq{"bt.id": filter.ID})
	}
	if filter.RoomID != "" {
		qb = qb.Where(sq.Eq{"bt.room_id": filter.RoomID})
	}
	if filter.UserID != "" {
		qb = qb.Where(sq.Eq{"bt.user_id": filter.UserID})
	}
	if filter.StartDate != "" || filter.EndDate != "" {
		qb = qb.Where(sq.Or{
			sq.Lt{"bt.start_date": filter.EndDate},
			sq.Gt{"bt.end_date": filter.StartDate},
		})
	}
	if filter.StartDate != "" {
		qb = qb.Where(sq.GtOrEq{"bt.start_date": filter.StartDate})
	}
	if filter.EndDate != "" {
		qb = qb.Where(sq.LtOrEq{"bt.end_date": filter.EndDate})
	}
	if filter.Status != "" {
		qb = qb.Where(sq.Eq{"bt.status": filter.Status})
	}

	return qb
}

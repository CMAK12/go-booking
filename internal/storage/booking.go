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
	ID        string `json:"id"`
	RoomID    string `json:"room_id"`
	UserID    string `json:"user_id"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Status    string `json:"status"`
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
			"bt.id", "bt.start_date", "bt.end_date", "bt.status",
			"rt.id", "rt.type", "rt.capacity", "rt.price", "rt.available",
			"ut.id", "ut.name", "ut.email", "ut.password", "ut.role", "ut.created_at",
			"ht.id", "ht.name", "ht.address", "ht.city", "ht.description", "ht.rating",
		).
		From(fmt.Sprintf("%s AS bt", bookingTable)).
		Join(fmt.Sprintf("%s AS rt ON bt.room_id = rt.id", roomTable)).
		Join(fmt.Sprintf("%s AS ut ON bt.user_id = ut.id", userTable)).
		Join(fmt.Sprintf("%s AS ht ON rt.hotel_id = ht.id", hotelTable))

	if filter.ID != "" {
		qb = qb.Where(sq.Eq{"bt.id": filter.ID})
	}
	if filter.RoomID != "" {
		qb = qb.Where(sq.Eq{"bt.room_id": filter.RoomID})
	}
	if filter.UserID != "" {
		qb = qb.Where(sq.Eq{"bt.user_id": filter.UserID})
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
			&booking.ID, &booking.StartDate, &booking.EndDate, &booking.Status,
			&booking.Room.ID, &booking.Room.Type, &booking.Room.Capacity, &booking.Room.Price, &booking.Room.Available,
			&booking.User.ID, &booking.User.Name, &booking.User.Email, &booking.User.Password, &booking.User.Role, &booking.User.CreatedAt,
			&booking.Room.Hotel.ID, &booking.Room.Hotel.Name, &booking.Room.Hotel.Address, &booking.Room.Hotel.City, &booking.Room.Hotel.Description, &booking.Room.Hotel.Rating,
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
		Values(booking.ID, booking.User.ID, booking.Room.ID, booking.StartDate, booking.EndDate, booking.Status).
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
		Set("user_id", booking.User.ID).
		Set("room_id", booking.Room.ID).
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

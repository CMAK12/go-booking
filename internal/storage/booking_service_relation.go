package storage

import (
	"context"
	"fmt"
	"go-booking/internal/dto"
	"go-booking/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type bookingServiceRelationStorage struct {
	db      *pgxpool.Pool
	builder sq.StatementBuilderType
}

type ListBookingServiceRelationFilter struct {
	BookingID      string `json:"booking_id"`
	ExtraServiceID string `json:"extra_id"`
}

func NewBookingServiceRelationStorage(db *pgxpool.Pool) BookingServiceRelationStorage {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &bookingServiceRelationStorage{
		db:      db,
		builder: builder,
	}
}

func (s *bookingServiceRelationStorage) List(ctx context.Context, filter ListBookingServiceRelationFilter) ([]models.BookingServiceRelation, error) {
	qb := s.builder.
		Select(
			"bs.booking_id", "bs.service_id",

			"b.start_date", "b.end_date", "b.status",

			"u.id AS user_id", "u.name AS user_name", "u.email AS user_email",
			"u.role AS user_role", "u.created_at AS user_created_at",

			"r.id AS room_id", "r.type AS room_type", "r.capacity AS room_capacity",
			"r.price AS room_price", "r.available AS room_available",

			"h.id AS hotel_id", "h.name AS hotel_name", "h.address AS hotel_address",
			"h.city AS hotel_city", "h.description AS hotel_description", "h.rating AS hotel_rating",

			"s.id AS service_id", "s.name AS service_name", "s.price AS service_price",
		).
		From(fmt.Sprintf("%s AS bs", bookingServiceRelationTable)).
		Join(fmt.Sprintf("%s b ON bs.booking_id = b.id", bookingTable)).
		Join(fmt.Sprintf("%s s ON bs.service_id = s.id", extraServiceTable)).
		Join(fmt.Sprintf("%s u ON b.user_id = u.id", userTable)).
		Join(fmt.Sprintf("%s r ON b.room_id = r.id", roomTable)).
		Join(fmt.Sprintf("%s h ON r.hotel_id = h.id", hotelTable))

	if filter.BookingID != "" {
		qb.Where(sq.Eq{"bs.booking_id": filter.BookingID})
	}
	if filter.ExtraServiceID != "" {
		qb.Where(sq.Eq{"bs.service_id": filter.ExtraServiceID})
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

	var relations []models.BookingServiceRelation

	for rows.Next() {
		var relation models.BookingServiceRelation

		err := rows.Scan(
			&relation.Booking.ID,
			&relation.Extra.ID,

			&relation.Booking.StartDate,
			&relation.Booking.EndDate,
			&relation.Booking.Status,

			&relation.Booking.User.ID,
			&relation.Booking.User.Name,
			&relation.Booking.User.Email,
			&relation.Booking.User.Role,
			&relation.Booking.User.CreatedAt,

			&relation.Booking.Room.ID,
			&relation.Booking.Room.Type,
			&relation.Booking.Room.Capacity,
			&relation.Booking.Room.Price,
			&relation.Booking.Room.Available,

			&relation.Booking.Room.Hotel.ID,
			&relation.Booking.Room.Hotel.Name,
			&relation.Booking.Room.Hotel.Address,
			&relation.Booking.Room.Hotel.City,
			&relation.Booking.Room.Hotel.Description,
			&relation.Booking.Room.Hotel.Rating,

			&relation.Extra.ID,
			&relation.Extra.Name,
			&relation.Extra.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		relations = append(relations, relation)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return relations, nil
}

func (s *bookingServiceRelationStorage) Create(ctx context.Context, bookindID, serviceID string) (dto.CreateBookingServiceRelationResponse, error) {
	query, args, err := s.builder.
		Insert(bookingServiceRelationTable).
		Columns("booking_id", "service_id").
		Values(bookindID, serviceID).
		Suffix("RETURNING booking_id, service_id").
		ToSql()
	if err != nil {
		return dto.CreateBookingServiceRelationResponse{}, fmt.Errorf("failed to build query: %w", err)
	}

	var bookingServiceRelation dto.CreateBookingServiceRelationResponse
	err = s.db.QueryRow(ctx, query, args...).Scan(&bookingServiceRelation.BookingID, &bookingServiceRelation.ServiceID)
	if err != nil {
		if pgx.ErrNoRows == err {
			return dto.CreateBookingServiceRelationResponse{}, fmt.Errorf("no rows found: %w", err)
		}
		return dto.CreateBookingServiceRelationResponse{}, fmt.Errorf("failed to create booking service relation: %w", err)
	}

	return bookingServiceRelation, nil
}

func (s *bookingServiceRelationStorage) Delete(ctx context.Context, bookingID, serviceID string) error {
	query, args, err := s.builder.
		Delete(bookingServiceRelationTable).
		Where(sq.And{
			sq.Eq{"booking_id": bookingID},
			sq.Eq{"service_id": serviceID},
		}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

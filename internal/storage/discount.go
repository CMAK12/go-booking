package storage

import (
	"context"
	"fmt"
	"go-booking/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type discountStorage struct {
	db      *pgxpool.Pool
	builder sq.StatementBuilderType
}

type ListDiscountFilter struct {
	ID      string  `json:"id"`
	HotelID string  `json:"hotel_id"`
	Amount  float64 `json:"amount"`
	Active  bool    `json:"active"`
}

func NewDiscountStorage(db *pgxpool.Pool) DiscountStorage {
	return &discountStorage{
		db:      db,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (s *discountStorage) List(ctx context.Context, filter ListDiscountFilter) ([]models.Discount, error) {
	qb := s.builder.
		Select(
			"d.id", "d.name", "d.amount", "d.active",
			"h.id", "h.name", "h.address", "h.city", "h.description", "h.rating",
		).
		From(fmt.Sprintf("%s AS d", discountTable)).
		Join(fmt.Sprintf("%s AS h ON d.hotel_id = h.id", hotelTable))

	if filter.ID != "" {
		qb = qb.Where(sq.Eq{"d.id": filter.ID})
	}
	if filter.HotelID != "" {
		qb = qb.Where(sq.Eq{"d.hotel_id": filter.HotelID})
	}
	if filter.Amount != 0 {
		qb = qb.Where(sq.Eq{"d.amount": filter.Amount})
	}
	if filter.Active {
		qb = qb.Where(sq.Eq{"d.active": filter.Active})
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

	var discounts []models.Discount
	for rows.Next() {
		var discount models.Discount
		err := rows.Scan(
			&discount.ID, &discount.Name, &discount.Amount, &discount.Active,
			&discount.Hotel.ID, &discount.Hotel.Name, &discount.Hotel.Address,
			&discount.Hotel.City, &discount.Hotel.Description, &discount.Hotel.Rating,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		discounts = append(discounts, discount)
	}

	return discounts, nil
}

func (s *discountStorage) Create(ctx context.Context, discount models.Discount) (models.Discount, error) {
	query, args, err := s.builder.
		Insert(discountTable).
		Columns("id", "hotel_id", "name", "amount", "active").
		Values(discount.ID, discount.Hotel.ID, discount.Name, discount.Amount, discount.Active).
		Suffix("RETURNING id, hotel_id, name, amount, active").
		ToSql()
	if err != nil {
		return models.Discount{}, fmt.Errorf("failed to build query: %w", err)
	}

	var createdDiscount models.Discount
	err = s.db.QueryRow(ctx, query, args...).Scan(&createdDiscount.ID, &createdDiscount.Hotel.ID, &createdDiscount.Name, &createdDiscount.Amount, &createdDiscount.Active)
	if err != nil {
		return models.Discount{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return createdDiscount, nil
}

func (s *discountStorage) Update(ctx context.Context, id string, discount models.Discount) (models.Discount, error) {
	query, args, err := s.builder.
		Update(discountTable).
		Set("name", discount.Name).
		Set("amount", discount.Amount).
		Set("active", discount.Active).
		Where(sq.Eq{"id": id}).
		Suffix("RETURNING id, hotel_id, name, amount, active").
		ToSql()
	if err != nil {
		return models.Discount{}, fmt.Errorf("failed to build query: %w", err)
	}

	var updatedDiscount models.Discount
	err = s.db.QueryRow(ctx, query, args...).Scan(&updatedDiscount.ID, &updatedDiscount.Hotel.ID, &updatedDiscount.Name, &updatedDiscount.Amount, &updatedDiscount.Active)
	if err != nil {
		return models.Discount{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return updatedDiscount, nil
}

func (s *discountStorage) Delete(ctx context.Context, id string) error {
	query, args, err := s.builder.
		Delete(discountTable).
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.Exec(ctx, query, args...)
	return err
}

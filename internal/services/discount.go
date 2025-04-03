package service

import (
	"context"
	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
)

type discountService struct {
	discountStorage storage.DiscountStorage
}

func NewDiscountService(discountStorage storage.DiscountStorage) DiscountService {
	return &discountService{
		discountStorage: discountStorage,
	}
}

func (s *discountService) List(ctx context.Context, filter storage.ListDiscountFilter) ([]models.Discount, error) {
	discounts, err := s.discountStorage.List(ctx, filter)
	if err != nil {
		return nil, err
	}
	return discounts, nil
}

func (s *discountService) Create(ctx context.Context, dt dto.CreateDiscountRequest) (models.Discount, error) {
	createdDiscount, err := s.discountStorage.Create(ctx, models.NewDiscount(dt))
	if err != nil {
		return models.Discount{}, err
	}
	return createdDiscount, nil
}

func (s *discountService) Update(ctx context.Context, id string, discount models.Discount) (models.Discount, error) {
	updatedDiscount, err := s.discountStorage.Update(ctx, id, discount)
	if err != nil {
		return models.Discount{}, err
	}
	return updatedDiscount, nil
}

func (s *discountService) Delete(ctx context.Context, id string) error {
	err := s.discountStorage.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

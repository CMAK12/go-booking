package service

import (
	"context"
	"fmt"
	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
)

type extraServiceService struct {
	extraServiceStorage storage.ExtraServiceStorage
}

func NewExtraServiceService(extraServiceStorage storage.ExtraServiceStorage) ExtraServiceService {
	return &extraServiceService{
		extraServiceStorage: extraServiceStorage,
	}
}

func (s *extraServiceService) List(ctx context.Context, filter storage.ListExtraServiceFilter) ([]models.ExtraService, error) {
	extraServices, err := s.extraServiceStorage.List(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to list extra services: %w", err)
	}
	return extraServices, nil
}

func (s *extraServiceService) Create(ctx context.Context, dto dto.CreateExtraServiceRequest) (models.ExtraService, error) {
	extraService := models.NewExtraService(
		dto.RoomID,
		dto.Name,
		dto.Price,
	)

	newExtraService, err := s.extraServiceStorage.Create(ctx, extraService)
	if err != nil {
		return models.ExtraService{}, fmt.Errorf("failed to create extra service: %w", err)
	}
	return newExtraService, nil
}

func (s *extraServiceService) Update(ctx context.Context, id string, extraService models.ExtraService) (models.ExtraService, error) {
	extraService, err := s.extraServiceStorage.Update(ctx, id, extraService)
	if err != nil {
		return models.ExtraService{}, fmt.Errorf("failed to update extra service: %w", err)
	}
	return extraService, nil
}

func (s *extraServiceService) Delete(ctx context.Context, id string) error {
	if err := s.extraServiceStorage.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete extra service: %w", err)
	}
	return nil
}

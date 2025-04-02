package service

import (
	"context"
	"log"

	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
)

type roomService struct {
	roomStorage storage.RoomStorage
}

func NewRoomService(roomStorage storage.RoomStorage) RoomService {
	return &roomService{
		roomStorage: roomStorage,
	}
}

func (s *roomService) Get(ctx context.Context, id string) (models.Room, error) {
	room, err := s.roomStorage.Get(ctx, id)
	if err != nil {
		log.Println("Error getting room:", err)
		return models.Room{}, err
	}

	return room, nil
}

func (s *roomService) List(ctx context.Context, filter storage.ListRoomFilter) ([]models.Room, error) {
	rooms, err := s.roomStorage.List(ctx, filter)
	if err != nil {
		log.Println("Error listing rooms:", err)
		return nil, err
	}
	return rooms, nil
}

func (s *roomService) Create(ctx context.Context, dto dto.CreateRoomRequest) (models.Room, error) {
	room, err := s.roomStorage.Create(ctx, models.NewRoom(dto))
	if err != nil {
		log.Println("Error creating room:", err)
		return models.Room{}, err
	}

	return room, nil
}

func (s *roomService) Update(ctx context.Context, id string, room models.Room) (models.Room, error) {
	room, err := s.roomStorage.Update(ctx, id, room)
	if err != nil {
		log.Println("Error updating room:", err)
		return models.Room{}, err
	}

	return room, nil
}

func (s *roomService) Delete(ctx context.Context, id string) error {
	err := s.roomStorage.Delete(ctx, id)
	if err != nil {
		log.Println("Error deleting room:", err)
		return err
	}

	return nil
}

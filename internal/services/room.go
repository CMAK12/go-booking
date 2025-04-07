package service

import (
	"context"
	"log"

	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
)

type roomService struct {
	roomStorage         storage.RoomStorage
	extraServiceStorage storage.ExtraServiceStorage
}

func NewRoomService(
	roomStorage storage.RoomStorage,
	extraServiceStorage storage.ExtraServiceStorage,
) RoomService {
	return &roomService{
		roomStorage:         roomStorage,
		extraServiceStorage: extraServiceStorage,
	}
}

func (s *roomService) List(ctx context.Context, filter storage.ListRoomFilter) ([]dto.ListRoomResponse, error) {
	rooms, err := s.roomStorage.List(ctx, filter)
	if err != nil {
		log.Println("Error listing rooms:", err)
		return nil, err
	}

	var responseRoom []dto.ListRoomResponse
	for _, room := range rooms {
		ess, err := s.extraServiceStorage.List(ctx, storage.ListExtraServiceFilter{RoomID: room.ID})
		if err != nil {
			log.Println("Error listing extra services:", err)
			return nil, err
		}

		responseRoom = append(responseRoom, dto.ListRoomResponse{
			ID:            room.ID,
			HotelID:       room.HotelID,
			ExtraServices: ess,
			Type:          room.Type,
			Capacity:      room.Capacity,
			Price:         room.Price,
			Quantity:      room.Quantity,
		})
	}

	return responseRoom, nil
}

func (s *roomService) Create(ctx context.Context, dto dto.CreateRoomRequest) (models.Room, error) {
	room := models.NewRoom(
		dto.HotelID,
		dto.Type,
		dto.Capacity,
		float64(dto.Price),
		dto.Quantity,
	)

	createdRoom, err := s.roomStorage.Create(ctx, room)
	if err != nil {
		log.Println("Error creating room:", err)
		return models.Room{}, err
	}

	createdRoom.HotelID = dto.HotelID

	return createdRoom, nil
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

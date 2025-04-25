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

func (s *roomService) List(ctx context.Context, rsFilter dto.ListRoomFilter) ([]dto.ListRoomResponse, int64, error) {
	rooms, rsCount, err := s.roomStorage.List(ctx, rsFilter)
	if err != nil {
		log.Println("Error listing rooms:", err)
		return nil, 0, err
	}

	if rsCount == 0 {
		return nil, 0, nil
	}

	roomIDs := make([]string, 0, rsCount)
	roomIDMap := make(map[string]bool, rsCount)

	for _, room := range rooms {
		if !roomIDMap[room.ID] {
			roomIDs = append(roomIDs, room.ID)
			roomIDMap[room.ID] = true
		}
	}

	extras, _, err := s.extraServiceStorage.List(ctx, dto.ListExtraServiceFilter{RoomIDs: roomIDs})
	if err != nil {
		log.Println("Error listing extra services:", err)
		return nil, 0, err
	}

	extraMap := make(map[string][]dto.ListExtraServiceResponse)
	for _, es := range extras {
		extraMap[es.RoomID] = append(extraMap[es.RoomID], dto.ListExtraServiceResponse{
			ID:    es.ID,
			Name:  es.Name,
			Price: es.Price,
		})
	}

	var responseRoom []dto.ListRoomResponse
	for _, room := range rooms {
		responseRoom = append(responseRoom, dto.ListRoomResponse{
			ID:            room.ID,
			HotelID:       room.HotelID,
			ExtraServices: extraMap[room.ID],
			Type:          room.Type,
			Capacity:      room.Capacity,
			Price:         room.Price,
			Quantity:      room.Quantity,
		})
	}

	return responseRoom, rsCount, nil
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

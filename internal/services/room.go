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
	hotelStorage        storage.HotelStorage
}

func NewRoomService(
	roomStorage storage.RoomStorage,
	extraServiceStorage storage.ExtraServiceStorage,
	hotelStorage storage.HotelStorage,
) RoomService {
	return &roomService{
		roomStorage:         roomStorage,
		extraServiceStorage: extraServiceStorage,
		hotelStorage:        hotelStorage,
	}
}

func (s *roomService) List(ctx context.Context, filter storage.ListRoomFilter) ([]models.Room, error) {
	rooms, err := s.roomStorage.List(ctx, filter)
	if err != nil {
		log.Println("Error listing rooms:", err)
		return nil, err
	}
	extraServices, err := s.extraServiceStorage.List(ctx, storage.ListExtraServiceFilter{
		RoomID: filter.ID,
	})
	if err != nil {
		log.Println("Error listing extra services in room service:", err)
		return nil, err
	}

	for i := range rooms {
		for _, extraService := range extraServices {
			if rooms[i].ID == extraService.RoomID {
				rooms[i].ExtraServices = append(rooms[i].ExtraServices, extraService)
			}
		}
	}

	return rooms, nil
}

func (s *roomService) Create(ctx context.Context, dto dto.CreateRoomRequest) (models.Room, error) {
	room, err := s.roomStorage.Create(ctx, models.NewRoom(dto))
	if err != nil {
		log.Println("Error creating room:", err)
		return models.Room{}, err
	}

	hotels, err := s.hotelStorage.List(ctx, storage.ListHotelFilter{
		ID: dto.HotelID,
	})
	if err != nil {
		log.Println("Error listing hotels in room service:", err)
		return models.Room{}, err
	}

	for i := range hotels {
		if hotels[i].ID == dto.HotelID {
			room.Hotel = hotels[i]
		}
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

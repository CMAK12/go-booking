package service

import (
	"context"
	"log"

	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"

	"github.com/redis/go-redis/v9"
)

type roomService struct {
	roomStorage         storage.RoomStorage
	extraServiceStorage storage.ExtraServiceStorage
	redisDB             *redis.Client
}

func NewRoomService(
	roomStorage storage.RoomStorage,
	extraServiceStorage storage.ExtraServiceStorage,
	redisDB *redis.Client,
) RoomService {
	return &roomService{
		roomStorage:         roomStorage,
		extraServiceStorage: extraServiceStorage,
		redisDB:             redisDB,
	}
}

const (
	keyRoomPopularitySet = "room:popularity:zset"
)

func (s *roomService) List(ctx context.Context, filter dto.ListRoomFilter) ([]dto.ListRoomResponse, int64, error) {
	popularRoomIDs, popularRoomIDsCount, err := s.getTopPopularRooms(ctx, filter.Skip, filter.Take)
	if err != nil {
		log.Println("Error getting top popular rooms:", err)
	}

	roomList := make([]models.Room, 0, filter.Take)

	if popularRoomIDsCount <= filter.Take {
		popularRooms, roomsCount, err := s.roomStorage.List(ctx, dto.ListRoomFilter{IDs: popularRoomIDs})
		if err != nil {
			log.Println("Error listing rooms by IDs:", err)
			return nil, 0, err
		}

		roomList = append(roomList, popularRooms...)

		if popularRoomIDsCount == filter.Take {
			roomResponses, err := s.buildRoomResponse(ctx, roomList)
			if err != nil {
				log.Println("Error building room response:", err)
				return nil, 0, err
			}

			return roomResponses, roomsCount, nil
		}
	}

	filter.Take = filter.Take - popularRoomIDsCount
	filter.ExcludeIDs = popularRoomIDs

	rooms, roomsCount, err := s.roomStorage.List(ctx, filter)
	if err != nil {
		log.Println("Error listing rooms:", err)
		return nil, 0, err
	}

	roomList = append(roomList, rooms...)

	go func(rooms []models.Room) {
		if err = s.incrementRoomPopularityScore(ctx, rooms); err != nil {
			log.Println("Error incrementing room popularity score:", err)
		}
	}(rooms)

	roomResponses, err := s.buildRoomResponse(ctx, roomList)
	if err != nil {
		log.Println("Error building room response:", err)
		return nil, 0, err
	}

	totalCount := popularRoomIDsCount + roomsCount

	return roomResponses, totalCount, nil
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

	err = s.removeRoomPopularityScore(ctx, id)
	if err != nil {
		log.Println("Error removing room from popularity set:", err)
	}

	return nil
}

func (s *roomService) buildRoomResponse(ctx context.Context, rooms []models.Room) ([]dto.ListRoomResponse, error) {
	result := make([]dto.ListRoomResponse, 0, len(rooms))

	for _, room := range rooms {
		extraServices, _, err := s.extraServiceStorage.List(ctx, storage.ListExtraServiceFilter{RoomID: room.ID})
		if err != nil {
			return nil, err
		}

		extraServiceResponses := make([]dto.ListExtraServiceResponse, len(extraServices))
		for i, es := range extraServices {
			extraServiceResponses[i] = dto.ListExtraServiceResponse{ID: es.ID, Name: es.Name, Price: es.Price}
		}

		result = append(result, dto.ListRoomResponse{
			ID:            room.ID,
			HotelID:       room.HotelID,
			ExtraServices: extraServiceResponses,
			Type:          room.Type,
			Capacity:      room.Capacity,
			Price:         room.Price,
			Quantity:      room.Quantity,
		})
	}

	return result, nil
}

func (s *roomService) incrementRoomPopularityScore(ctx context.Context, rooms []models.Room) error {
	for _, room := range rooms {
		err := s.redisDB.ZIncrBy(ctx, keyRoomPopularitySet, 1, room.ID).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *roomService) getTopPopularRooms(ctx context.Context, start int64, stop int64) ([]string, int64, error) {
	popularRooms, err := s.redisDB.ZRevRange(ctx, keyRoomPopularitySet, start, stop-1).Result()
	if err != nil {
		return nil, 0, err
	}

	popularRoomsCount := int64(len(popularRooms))

	return popularRooms, popularRoomsCount, nil
}

func (s *roomService) removeRoomPopularityScore(ctx context.Context, roomID string) error {
	err := s.redisDB.ZRem(ctx, keyRoomPopularitySet, roomID).Err()
	if err != nil {
		return err
	}

	return nil
}

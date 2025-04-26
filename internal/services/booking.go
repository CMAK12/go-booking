package service

import (
	"context"
	"fmt"
	"log"
	"math"
	"sort"

	"go-booking/internal/broker"
	"go-booking/internal/consts"
	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
)

type bookingService struct {
	emailBroker         *broker.EmailBroker
	bookingStorage      storage.BookingStorage
	hotelStorage        storage.HotelStorage
	roomStorage         storage.RoomStorage
	roomService         RoomService
	userStorage         storage.UserStorage
	extraServiceStorage storage.ExtraServiceStorage
}

func NewBookingService(
	emailBroker *broker.EmailBroker,
	bookingStorage storage.BookingStorage,
	hotelStorage storage.HotelStorage,
	roomStorage storage.RoomStorage,
	roomService RoomService,
	userStorage storage.UserStorage,
	extraServiceStorage storage.ExtraServiceStorage,
) BookingService {
	return &bookingService{
		emailBroker:         emailBroker,
		bookingStorage:      bookingStorage,
		hotelStorage:        hotelStorage,
		roomStorage:         roomStorage,
		roomService:         roomService,
		userStorage:         userStorage,
		extraServiceStorage: extraServiceStorage,
	}
}

const (
	dateLayout       = "2006-01-02"
	hoursInDay       = 24
	errorMsgTemplate = "room %s is already booked for the selected dates. Nearest free dates: %s"
	storageErrorMsg  = "failed to check booking availability: %w"
)

func (s *bookingService) List(ctx context.Context, filter dto.ListBookingFilter) ([]dto.ListBookingResponse, int64, error) {
	bookings, count, err := s.bookingStorage.List(ctx, filter)
	if err != nil {
		log.Println("failed to list bookings:", err)
		return nil, 0, err
	}

	if count == 0 {
		return []dto.ListBookingResponse{}, 0, nil
	}

	userIDs := make([]string, 0, count)
	roomIDs := make([]string, 0, count)
	userIDMap := make(map[string]bool)
	roomIDMap := make(map[string]bool)

	for _, booking := range bookings {
		if !userIDMap[booking.UserID] {
			userIDs = append(userIDs, booking.UserID)
			userIDMap[booking.UserID] = true
		}

		if !roomIDMap[booking.RoomID] {
			roomIDs = append(roomIDs, booking.RoomID)
			roomIDMap[booking.RoomID] = true
		}
	}

	users, _, err := s.userStorage.List(ctx, dto.ListUserFilter{IDs: userIDs})
	if err != nil {
		log.Println("failed to list users:", err)
		return nil, 0, err
	}
	userMap := make(map[string]models.User)
	for _, user := range users {
		userMap[user.ID] = user
	}

	rooms, roomsCount, err := s.roomService.List(ctx, dto.ListRoomFilter{IDs: roomIDs})
	if err != nil {
		log.Println("failed to list rooms:", err)
		return nil, 0, err
	}
	roomMap := make(map[string]dto.ListRoomResponse)

	hotelIDs := make([]string, 0, roomsCount)
	hotelIDMap := make(map[string]bool)
	for _, room := range rooms {
		if !hotelIDMap[room.HotelID] {
			hotelIDs = append(hotelIDs, room.HotelID)
			hotelIDMap[room.HotelID] = true
		}
		roomMap[room.ID] = room
	}

	hotels, _, err := s.hotelStorage.List(ctx, dto.ListHotelFilter{IDs: hotelIDs})
	if err != nil {
		log.Println("failed to list hotels:", err)
		return nil, 0, err
	}
	hotelMap := make(map[string]models.Hotel)
	for _, hotel := range hotels {
		hotelMap[hotel.ID] = hotel
	}

	bookingResponse := make([]dto.ListBookingResponse, 0, count)
	for _, booking := range bookings {
		user, userExists := userMap[booking.UserID]
		if !userExists {
			log.Println("user not found for booking:", booking.ID)
			continue
		}

		room, roomExists := roomMap[booking.RoomID]
		if !roomExists {
			log.Println("room not found for booking:", booking.ID)
			continue
		}

		hotel, hotelExists := hotelMap[room.HotelID]
		if !hotelExists {
			log.Println("hotel not found for booking:", booking.ID)
			continue
		}

		response := dto.ListBookingResponse{
			ID:        booking.ID,
			User:      user,
			Room:      room,
			Hotel:     hotel,
			StartDate: booking.StartDate,
			EndDate:   booking.EndDate,
			Status:    string(booking.Status),
		}
		bookingResponse = append(bookingResponse, response)
	}

	return bookingResponse, count, nil
}

func (s *bookingService) Create(ctx context.Context, bookingDTO dto.CreateBookingRequest) (models.Booking, error) {
	booking, err := models.NewBooking(
		bookingDTO.UserID,
		bookingDTO.RoomID,
		bookingDTO.HotelID,
		bookingDTO.StartDate,
		bookingDTO.EndDate,
	)
	if err != nil {
		log.Printf("failed to create booking: %v", err)
		return models.Booking{}, err
	}

	if err := s.checkRoomAvailability(ctx, booking, bookingDTO); err != nil {
		return models.Booking{}, err
	}

	newBooking, err := s.bookingStorage.Create(ctx, booking)
	if err != nil {
		log.Printf("failed to store booking: %v", err)
		return models.Booking{}, err
	}

	users, userCount, err := s.userStorage.List(ctx, dto.ListUserFilter{ID: booking.UserID})
	if err != nil || userCount == 0 {
		log.Printf("failed to fetch user for booking email: %v", err)
		return newBooking, nil
	}
	user := users[0]

	s.emailBroker.Publish(broker.EmailTask{
		To:      user.Email,
		Subject: consts.BookingVerificationSubject,
		Body:    consts.BookingVerificationBody,
	})

	return newBooking, nil
}

func (s *bookingService) Update(ctx context.Context, id string, dto dto.UpdateBookingRequest) (models.Booking, error) {
	booking, err := models.NewBookingFromDTO(
		id,
		dto.UserID,
		dto.RoomID,
		dto.StartDate,
		dto.EndDate,
		dto.Status,
	)
	if err != nil {
		log.Println("failed to create booking from DTO:", err)
		return models.Booking{}, err
	}

	updatedBooking, err := s.bookingStorage.Update(ctx, id, booking)
	if err != nil {
		log.Println("failed to update booking:", err)
		return models.Booking{}, err
	}

	return updatedBooking, nil
}

func (s *bookingService) Delete(ctx context.Context, id string) error {
	err := s.bookingStorage.Delete(ctx, id)
	if err != nil {
		log.Println("failed to delete booking:", err)
		return err
	}
	return nil
}

func (s *bookingService) checkRoomAvailability(ctx context.Context, booking models.Booking, bookingDTO dto.CreateBookingRequest) error {
	overlapBookings, overlapCount, err := s.bookingStorage.List(ctx, dto.ListBookingFilter{
		RoomID:    booking.RoomID,
		StartDate: booking.StartDate,
		EndDate:   booking.EndDate,
	})
	if err != nil {
		log.Printf("failed to check room availability: %v", err)
		return err
	}

	rooms, roomsCount, err := s.roomStorage.List(ctx, dto.ListRoomFilter{ID: booking.RoomID})
	if err != nil || roomsCount == 0 {
		log.Printf("failed to retrieve room for booking: %v", err)
		return err
	}
	room := rooms[0]

	isAvailable := int64(room.Quantity) <= overlapCount
	if overlapCount > 0 && isAvailable {
		return s.findNextAvailableDate(ctx, booking, bookingDTO, overlapBookings[0])
	}

	return nil
}

func (s *bookingService) findNextAvailableDate(ctx context.Context, booking models.Booking, bookingDTO dto.CreateBookingRequest, overlapBooking models.Booking) error {
	filter := dto.ListBookingFilter{
		RoomID:     booking.RoomID,
		LatestDate: bookingDTO.EndDate,
		Status:     []models.BookingStatus{models.BookingStatusPending, models.BookingStatusConfirmed},
	}

	bookings, count, err := s.bookingStorage.List(ctx, filter)
	if err != nil {
		return fmt.Errorf(storageErrorMsg, err)
	}

	duration := booking.EndDate.Sub(booking.StartDate)

	var nextAvailableDateRange string

	if count > 0 {
		sort.Slice(bookings, func(i, j int) bool {
			return bookings[i].StartDate.Before(bookings[j].StartDate)
		})

		for i := int64(0); i < count-1; i++ {
			current := bookings[i]
			next := bookings[i+1]

			if next.StartDate.Sub(current.EndDate) >= duration {
				nextAvailableDateRange = fmt.Sprintf("%s - %s",
					current.EndDate.Format(dateLayout),
					current.EndDate.Add(duration).Format(dateLayout))
				break
			}
		}

		if nextAvailableDateRange == "" {
			lastBooking := bookings[count-1]
			nextAvailableDateRange = fmt.Sprintf("%s - %s",
				lastBooking.EndDate.Format(dateLayout),
				lastBooking.EndDate.Add(duration).Format(dateLayout))
		}
	} else {
		daysToMove := int(math.Round(overlapBooking.EndDate.Sub(booking.StartDate).Hours() / hoursInDay))
		nearestStart := booking.StartDate.AddDate(0, 0, daysToMove)
		nextAvailableDateRange = fmt.Sprintf("%s - %s",
			nearestStart.Format(dateLayout),
			nearestStart.Add(duration).Format(dateLayout))
	}

	return fmt.Errorf(errorMsgTemplate, booking.RoomID, nextAvailableDateRange)
}

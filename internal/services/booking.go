package service

import (
	"context"
	"fmt"
	"log"

	"go-booking/internal/consts"
	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"

	"math"

	"github.com/veyselaksin/gomailer/pkg/mailer"
)

type bookingService struct {
	bookingStorage      storage.BookingStorage
	hotelStorage        storage.HotelStorage
	roomStorage         storage.RoomStorage
	roomService         RoomService
	userStorage         storage.UserStorage
	extraServiceStorage storage.ExtraServiceStorage
	mailAuth            mailer.Authentication
}

func NewBookingService(
	bookingStorage storage.BookingStorage,
	hotelStorage storage.HotelStorage,
	roomStorage storage.RoomStorage,
	roomService RoomService,
	userStorage storage.UserStorage,
	extraServiceStorage storage.ExtraServiceStorage,
	mailAuth mailer.Authentication,
) BookingService {
	return &bookingService{
		bookingStorage:      bookingStorage,
		hotelStorage:        hotelStorage,
		roomStorage:         roomStorage,
		roomService:         roomService,
		userStorage:         userStorage,
		extraServiceStorage: extraServiceStorage,
		mailAuth:            mailAuth,
	}
}

func (s *bookingService) List(ctx context.Context, filter storage.ListBookingFilter) ([]dto.ListBookingResponse, int64, error) {
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

	users, _, err := s.userStorage.List(ctx, storage.ListUserFilter{IDs: userIDs})
	if err != nil {
		log.Println("failed to list users:", err)
		return nil, 0, err
	}
	userMap := make(map[string]models.User)
	for _, user := range users {
		userMap[user.ID] = user
	}

	rooms, roomsCount, err := s.roomService.List(ctx, storage.ListRoomFilter{IDs: roomIDs})
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

	hotels, _, err := s.hotelStorage.List(ctx, storage.ListHotelFilter{IDs: hotelIDs})
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

func (s *bookingService) Create(ctx context.Context, dto dto.CreateBookingRequest) (models.Booking, error) {
	booking, err := models.NewBooking(
		dto.UserID,
		dto.RoomID,
		dto.HotelID,
		dto.StartDate,
		dto.EndDate,
	)
	if err != nil {
		log.Printf("failed to create booking: %v", err)
		return models.Booking{}, err
	}

	if err := s.checkRoomAvailability(ctx, booking, dto); err != nil {
		return models.Booking{}, err
	}

	newBooking, err := s.bookingStorage.Create(ctx, booking)
	if err != nil {
		log.Printf("failed to store booking: %v", err)
		return models.Booking{}, err
	}

	users, userCount, err := s.userStorage.List(ctx, storage.ListUserFilter{ID: booking.UserID})
	if err != nil || userCount == 0 {
		log.Printf("failed to fetch user for booking email: %v", err)
		return newBooking, nil
	}
	user := users[0]

	go func(userEmail string) {
		sender := mailer.NewPlainAuth(&s.mailAuth)
		msg := mailer.NewMessage(consts.BookingVerificationSubject, consts.BookingVerificationBody)
		msg.SetTo([]string{userEmail})

		if err := sender.SendMail(msg); err != nil {
			log.Printf("failed to send email to %s: %v", userEmail, err)
			return
		}
		log.Printf("email sent successfully to: %s", userEmail)
	}(user.Email)

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

func (s *bookingService) checkRoomAvailability(ctx context.Context, booking models.Booking, dto dto.CreateBookingRequest) error {
	overlapBookings, overlapCount, err := s.bookingStorage.List(ctx, storage.ListBookingFilter{
		RoomID:    booking.RoomID,
		StartDate: booking.StartDate,
		EndDate:   booking.EndDate,
	})
	if err != nil {
		log.Printf("failed to check room availability: %v", err)
		return err
	}

	rooms, roomsCount, err := s.roomStorage.List(ctx, storage.ListRoomFilter{ID: booking.RoomID})
	if err != nil || roomsCount == 0 {
		log.Printf("failed to retrieve room for booking: %v", err)
		return err
	}
	room := rooms[0]

	isAvailable := int64(room.Quantity) <= overlapCount
	if overlapCount > 0 && isAvailable {
		return s.findNextAvailableDates(ctx, booking, dto, overlapBookings[0])
	}

	return nil
}

func (s *bookingService) findNextAvailableDates(ctx context.Context, booking models.Booking, dto dto.CreateBookingRequest, overlapBooking models.Booking) error {
	bookings, count, err := s.bookingStorage.List(ctx, storage.ListBookingFilter{
		RoomID:     booking.RoomID,
		LatestDate: dto.EndDate,
		Status:     []models.BookingStatus{models.BookingStatusPending, models.BookingStatusConfirmed},
	})
	if err != nil {
		log.Printf("failed to retrieve conflicting bookings: %v", err)
		return err
	}

	var nextAvailableDateRange string
	if count > 0 {
		mostRecentBooking := bookings[0]
		duration := booking.EndDate.Sub(booking.StartDate).Hours()

		for _, currentBooking := range bookings[1:] {
			gapToNextBookingHours := currentBooking.StartDate.Sub(mostRecentBooking.EndDate).Hours()

			if gapToNextBookingHours >= duration {
				break
			}

			mostRecentBooking = currentBooking
		}

		days := duration / 24
		nextAvailableEndDate := mostRecentBooking.EndDate.AddDate(0, 0, int(math.Round(days)))

		nextAvailableDateRange = fmt.Sprintf("%s - %s",
			mostRecentBooking.EndDate.Format("2006-01-02"),
			nextAvailableEndDate.Format("2006-01-02"))
	} else {
		days := math.Round(overlapBooking.EndDate.Sub(booking.StartDate).Hours() / 24)
		nearestStart := booking.StartDate.AddDate(0, 0, int(days))
		nearestEnd := booking.EndDate.AddDate(0, 0, int(days))

		nextAvailableDateRange = fmt.Sprintf("%s - %s",
			nearestStart.Format("2006-01-02"),
			nearestEnd.Format("2006-01-02"))
	}

	const msgTemplate = "room %s is already booked for the selected dates. Nearest free dates: %s"
	msg := fmt.Sprintf(msgTemplate, booking.RoomID, nextAvailableDateRange)

	log.Println(msg)
	return fmt.Errorf("%s", msg)
}

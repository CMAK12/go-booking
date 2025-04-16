package service

import (
	"context"
	"fmt"
	"log"

	"go-booking/internal/consts"
	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"

	"github.com/veyselaksin/gomailer/pkg/mailer"
)

type bookingService struct {
	bookingStorage      storage.BookingStorage
	hotelStorage        storage.HotelStorage
	roomService         RoomService
	userStorage         storage.UserStorage
	extraServiceStorage storage.ExtraServiceStorage
	mailAuth            mailer.Authentication
}

func NewBookingService(
	bookingStorage storage.BookingStorage,
	hotelStorage storage.HotelStorage,
	roomService RoomService,
	userStorage storage.UserStorage,
	extraServiceStorage storage.ExtraServiceStorage,
	mailAuth mailer.Authentication,
) BookingService {
	return &bookingService{
		bookingStorage:      bookingStorage,
		hotelStorage:        hotelStorage,
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

	rooms, rsCount, err := s.roomService.List(ctx, storage.ListRoomFilter{IDs: roomIDs})
	if err != nil {
		log.Println("failed to list rooms:", err)
		return nil, 0, err
	}
	roomMap := make(map[string]dto.ListRoomResponse)

	hotelIDs := make([]string, 0, rsCount)
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

	// Check room availability
	_, overlapCount, err := s.bookingStorage.List(ctx, storage.ListBookingFilter{
		RoomID:    dto.RoomID,
		StartDate: dto.StartDate,
		EndDate:   dto.EndDate,
	})
	if err != nil {
		log.Printf("failed to check room availability: %v", err)
		return models.Booking{}, err
	}

	if overlapCount > 0 {
		bookings, count, err := s.bookingStorage.List(ctx, storage.ListBookingFilter{
			RoomID:     dto.RoomID,
			LatestDate: dto.EndDate,
			Status:     []models.BookingStatus{models.BookingStatusPending, models.BookingStatusConfirmed},
		})
		if err != nil {
			log.Printf("failed to retrieve conflicting bookings: %v", err)
			return models.Booking{}, err
		}

		var nearestFreeDate string
		if count > 0 {
			latestEndDate := bookings[0].EndDate
			for _, b := range bookings[1:] {
				if b.EndDate.After(latestEndDate) {
					latestEndDate = b.EndDate
				}
			}

			days := int(booking.EndDate.Sub(booking.StartDate).Hours() / 24)
			nextAvailableStart := latestEndDate.AddDate(0, 0, 1)
			nextAvailableEnd := nextAvailableStart.AddDate(0, 0, days)

			nearestFreeDate = fmt.Sprintf("%s - %s",
				nextAvailableStart.Format("2006-01-02"),
				nextAvailableEnd.Format("2006-01-02"))
		} else {
			days := int(booking.EndDate.Sub(booking.StartDate).Hours() / 24)
			nearestEnd := booking.StartDate.AddDate(0, 0, days)
			nearestFreeDate = fmt.Sprintf("%s - %s",
				booking.StartDate.Format("2006-01-02"),
				nearestEnd.Format("2006-01-02"))
		}

		msg := fmt.Sprintf("room %s is already booked for the selected dates. Nearest free dates: %s", booking.RoomID, nearestFreeDate)
		log.Println(msg)
		return models.Booking{}, fmt.Errorf("%s", msg)
	}

	newBooking, err := s.bookingStorage.Create(ctx, booking)
	if err != nil {
		log.Printf("failed to store booking: %v", err)
		return models.Booking{}, err
	}

	users, usCount, err := s.userStorage.List(ctx, storage.ListUserFilter{ID: booking.UserID})
	if err != nil || usCount == 0 {
		log.Printf("failed to fetch user for booking email: %v", err)
		return newBooking, nil
	}

	go func(userEmail string) {
		sender := mailer.NewPlainAuth(&s.mailAuth)
		msg := mailer.NewMessage(consts.BookingVerificationSubject, consts.BookingVerificationBody)
		msg.SetTo([]string{userEmail})

		if err := sender.SendMail(msg); err != nil {
			log.Printf("failed to send email to %s: %v", userEmail, err)
			return
		}
		log.Printf("email sent successfully to: %s", userEmail)
	}(users[0].Email)

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

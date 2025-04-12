package service

import (
	"context"
	"log"

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

func (s *bookingService) List(ctx context.Context, filter storage.ListBookingFilter) ([]dto.ListBookingResponse, error) {
	bookings, err := s.bookingStorage.List(ctx, filter)
	if err != nil {
		log.Println("failed to list bookings:", err)
		return nil, err
	}

	if len(bookings) == 0 {
		return []dto.ListBookingResponse{}, nil
	}

	userIDs := make([]string, 0, len(bookings))
	roomIDs := make([]string, 0, len(bookings))
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

	users, err := s.userStorage.List(ctx, storage.ListUserFilter{IDs: userIDs})
	if err != nil {
		log.Println("failed to list users:", err)
		return nil, err
	}
	userMap := make(map[string]models.User)
	for _, user := range users {
		userMap[user.ID] = user
	}

	rooms, err := s.roomService.List(ctx, storage.ListRoomFilter{IDs: roomIDs})
	if err != nil {
		log.Println("failed to list rooms:", err)
		return nil, err
	}
	roomMap := make(map[string]dto.ListRoomResponse)

	hotelIDs := make([]string, 0, len(rooms))
	hotelIDMap := make(map[string]bool)
	for _, room := range rooms {
		if !hotelIDMap[room.HotelID] {
			hotelIDs = append(hotelIDs, room.HotelID)
			hotelIDMap[room.HotelID] = true
		}
		roomMap[room.ID] = room
	}

	hotels, err := s.hotelStorage.List(ctx, storage.ListHotelFilter{IDs: hotelIDs})
	if err != nil {
		log.Println("failed to list hotels:", err)
		return nil, err
	}
	hotelMap := make(map[string]models.Hotel)
	for _, hotel := range hotels {
		hotelMap[hotel.ID] = hotel
	}

	bookingResponse := make([]dto.ListBookingResponse, 0, len(bookings))
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

	return bookingResponse, nil
}

func (s *bookingService) Create(ctx context.Context, dto dto.CreateBookingRequest) (models.Booking, error) {
	booking := models.NewBooking(
		dto.UserID,
		dto.RoomID,
		dto.HotelID,
		dto.StartDate,
		dto.EndDate,
	)

	newBooking, err := s.bookingStorage.Create(ctx, booking)
	if err != nil {
		log.Println("failed to create booking:", err)
		return models.Booking{}, err
	}

	// go func() {
	// 	sender := mailer.NewPlainAuth(&s.mailAuth)
	// 	message := mailer.NewMessage(
	// 		consts.BookingVerificationSubject,
	// 		consts.BookingVerificationBody,
	// 	)
	// 	message.SetTo([]string{booking.User.Email})

	// 	if err := sender.SendMail(message); err != nil {
	// 		log.Printf("failed to send email to %s: %v", booking.User.Email, err)
	// 		return
	// 	}

	// 	log.Println("email sent successfully to:", booking.User.Email)
	// }()

	bookings, err := s.List(ctx, storage.ListBookingFilter{ID: booking.ID})
	if err != nil {
		log.Println("failed to list bookings after creation:", err)
		return models.Booking{}, err
	}

	for _, b := range bookings {
		if b.ID == booking.ID {
			bookings[0] = b
			break
		}
	}

	// hotels, err := s.hotelStorage.List(ctx, storage.ListHotelFilter{ID: booking.Room.Hotel.ID})
	// if err != nil {
	// 	log.Println("failed to list hotels after booking creation:", err)
	// 	return models.Booking{}, err
	// }

	return newBooking, nil
}

func (s *bookingService) Update(ctx context.Context, id string, dto dto.UpdateBookingRequest) (models.Booking, error) {
	booking := models.NewBookingFromDTO(
		id,
		dto.UserID,
		dto.RoomID,
		dto.StartDate,
		dto.EndDate,
		dto.Status,
	)

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

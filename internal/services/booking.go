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

	var bookingResponse []dto.ListBookingResponse
	for _, booking := range bookings {
		users, err := s.userStorage.List(ctx, storage.ListUserFilter{ID: booking.UserID})
		if err != nil || len(users) == 0 {
			log.Println("failed to list users:", err)
			return nil, err
		}
		user := users[0]

		rooms, err := s.roomService.List(ctx, storage.ListRoomFilter{ID: booking.RoomID})
		if err != nil || len(rooms) == 0 {
			log.Println("failed to list rooms:", err)
			return nil, err
		}
		room := rooms[0]

		hotels, err := s.hotelStorage.List(ctx, storage.ListHotelFilter{ID: room.HotelID})
		if err != nil || len(hotels) == 0 {
			log.Println("failed to list hotels:", err)
			return nil, err
		}
		hotel := hotels[0]

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

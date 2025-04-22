package v1

import (
	service "go-booking/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	userService    service.UserService
	bookingService service.BookingService
	hotelService   service.HotelService
	roomService    service.RoomService
	extraService   service.ExtraServiceService
}

func NewHandler(
	userService service.UserService,
	bookingService service.BookingService,
	hotelService service.HotelService,
	roomService service.RoomService,
	extraService service.ExtraServiceService,
) *Handler {
	return &Handler{
		userService:    userService,
		bookingService: bookingService,
		hotelService:   hotelService,
		roomService:    roomService,
		extraService:   extraService,
	}
}

func (h *Handler) SetupRoutes(r *chi.Mux) {
	r.Use(Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", ResponseWrapper(h.listUser))
			r.Post("/", ResponseWrapper(h.createUser))
			r.Put("/{id}", ResponseWrapper(h.updateUser))
			r.Delete("/{id}", ResponseWrapper(h.deleteUser))
		})
		r.Route("/bookings", func(r chi.Router) {
			r.Get("/", ResponseWrapper(h.listBooking))
			r.Post("/", ResponseWrapper(h.createBooking))
			r.Put("/{id}", ResponseWrapper(h.updateBooking))
			r.Delete("/{id}", ResponseWrapper(h.deleteBooking))
		})
		r.Route("/hotels", func(r chi.Router) {
			r.Get("/", ResponseWrapper(h.listHotel))
			r.Post("/", ResponseWrapper(h.createHotel))
			r.Put("/{id}", ResponseWrapper(h.updateHotel))
			r.Delete("/{id}", ResponseWrapper(h.deleteHotel))
		})
		r.Route("/rooms", func(r chi.Router) {
			r.Get("/", ResponseWrapper(h.listRoom))
			r.Post("/", ResponseWrapper(h.createRoom))
			r.Put("/{id}", ResponseWrapper(h.updateRoom))
			r.Delete("/{id}", ResponseWrapper(h.deleteRoom))
		})
		r.Route("/services", func(r chi.Router) {
			r.Get("/", ResponseWrapper(h.listExtraService))
			r.Post("/", ResponseWrapper(h.createExtraService))
			r.Put("/{id}", ResponseWrapper(h.updateExtraService))
			r.Delete("/{id}", ResponseWrapper(h.deleteExtraService))
		})
	})
}

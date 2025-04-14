package v1

import (
	"encoding/json"
	"net/http"

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
			r.Get("/", h.listUser)
			r.Post("/", h.createUser)
			r.Put("/{id}", h.updateUser)
			r.Delete("/{id}", h.deleteUser)
		})
		r.Route("/bookings", func(r chi.Router) {
			r.Get("/", h.listBooking)
			r.Post("/", h.createBooking)
			r.Put("/{id}", h.updateBooking)
			r.Delete("/{id}", h.deleteBooking)
		})
		r.Route("/hotels", func(r chi.Router) {
			r.Get("/", h.listHotel)
			r.Post("/", h.createHotel)
			r.Put("/{id}", h.updateHotel)
			r.Delete("/{id}", h.deleteHotel)
		})
		r.Route("/rooms", func(r chi.Router) {
			r.Get("/", h.listRoom)
			r.Post("/", h.createRoom)
			r.Put("/{id}", h.updateRoom)
			r.Delete("/{id}", h.deleteRoom)
		})
		r.Route("/services", func(r chi.Router) {
			r.Get("/", h.listExtraService)
			r.Post("/", h.createExtraService)
			r.Put("/{id}", h.updateExtraService)
			r.Delete("/{id}", h.deleteExtraService)
		})
	})
}

func writeJSON(w http.ResponseWriter, status int, data interface{}, count ...int64) {
	response := map[string]interface{}{
		"data": data,
	}

	if len(count) > 0 {
		response["count"] = count[0]
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}

package v1

import (
	service "go-booking/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	userService        service.UserService
	reservationService service.ReservationService
}

func NewHandler(
		userService service.UserService,
		reservationService service.ReservationService,
	) *Handler {
	return &Handler{
		userService:        userService,
		reservationService: reservationService,
	}
}

func (h *Handler) SetupRoutes(r *chi.Mux) {
	r.Use(Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", h.listUser)
			r.Get("/{id}", h.getUser)
			r.Post("/", h.createUser)
			r.Put("/{id}", h.updateUser)
			r.Delete("/{id}", h.deleteUser)
    })
		r.Route("/reservations", func(r chi.Router) {
			r.Get("/", nil)
			// r.Post("/", CreateReservation)
			// r.Get("/{id}", GetReservation)
		})
	})
}

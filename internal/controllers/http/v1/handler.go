package v1

import (
	service "go-booking/internal/services"

	"github.com/gofiber/fiber/v2"
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

func (h *Handler) SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1", Logger())

	users := api.Group("/users")
	users.Get("/", ResponseWrapper(h.listUser))
	users.Post("/", ResponseWrapper(h.createUser))
	users.Put("/:id", ResponseWrapper(h.updateUser))
	users.Delete("/:id", ResponseWrapper(h.deleteUser))

	bookings := api.Group("/bookings")
	bookings.Get("/", ResponseWrapper(h.listBooking))
	bookings.Post("/", ResponseWrapper(h.createBooking))
	bookings.Put("/:id", ResponseWrapper(h.updateBooking))
	bookings.Delete("/:id", ResponseWrapper(h.deleteBooking))

	hotels := api.Group("/hotels")
	hotels.Get("/", ResponseWrapper(h.listHotel))
	hotels.Post("/", ResponseWrapper(h.createHotel))
	hotels.Put("/:id", ResponseWrapper(h.updateHotel))
	hotels.Delete("/:id", ResponseWrapper(h.deleteHotel))

	rooms := api.Group("/rooms")
	rooms.Get("/", ResponseWrapper(h.listRoom))
	rooms.Post("/", ResponseWrapper(h.createRoom))
	rooms.Put("/:id", ResponseWrapper(h.updateRoom))
	rooms.Delete("/:id", ResponseWrapper(h.deleteRoom))

	services := api.Group("/services")
	services.Get("/", ResponseWrapper(h.listExtraService))
	services.Post("/", ResponseWrapper(h.createExtraService))
	services.Put("/:id", ResponseWrapper(h.updateExtraService))
	services.Delete("/:id", ResponseWrapper(h.deleteExtraService))
}

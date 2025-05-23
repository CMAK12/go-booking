package v1

import (
	"go-booking/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) listBooking(c *fiber.Ctx) (any, int, int64, error) {
	var filter dto.ListBookingFilter

	if err := c.QueryParser(&filter); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	bookings, count, err := h.bookingService.List(c.Context(), filter)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return bookings, fiber.StatusOK, count, nil
}

func (h *Handler) createBooking(c *fiber.Ctx) (any, int, int64, error) {
	var dto dto.CreateBookingRequest
	if err := c.BodyParser(&dto); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	booking, err := h.bookingService.Create(c.Context(), dto)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return booking, fiber.StatusCreated, 1, nil
}

func (h *Handler) updateBooking(c *fiber.Ctx) (any, int, int64, error) {
	id := c.Params("id")

	var booking dto.UpdateBookingRequest
	if err := c.BodyParser(&booking); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	updatedBooking, err := h.bookingService.Update(c.Context(), id, booking)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return updatedBooking, fiber.StatusOK, 1, nil
}

func (h *Handler) deleteBooking(c *fiber.Ctx) (any, int, int64, error) {
	id := c.Params("id")

	if err := h.bookingService.Delete(c.Context(), id); err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return nil, fiber.StatusNoContent, 0, nil
}

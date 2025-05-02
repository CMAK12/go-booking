package v1

import (
	"go-booking/internal/dto"
	"go-booking/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) listHotel(c *fiber.Ctx) (any, int, int64, error) {
	var filter dto.ListHotelFilter

	if err := c.QueryParser(&filter); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	if rating := c.Query("rating"); rating != "" {
		ratingFloat, err := strconv.ParseFloat(rating, 64)
		if err != nil {
			return nil, fiber.StatusBadRequest, 0, err
		}
		filter.Rating = ratingFloat
	}

	hotels, count, err := h.hotelService.List(c.Context(), filter)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return hotels, fiber.StatusOK, count, nil
}

func (h *Handler) createHotel(c *fiber.Ctx) (any, int, int64, error) {
	var dto dto.CreateHotelRequest
	if err := c.BodyParser(&dto); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	hotel, err := h.hotelService.Create(c.Context(), dto)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return hotel, fiber.StatusCreated, 1, nil
}

func (h *Handler) updateHotel(c *fiber.Ctx) (any, int, int64, error) {
	id := c.Params("id")

	var hotel models.Hotel
	if err := c.BodyParser(&hotel); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	updatedHotel, err := h.hotelService.Update(c.Context(), id, hotel)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return updatedHotel, fiber.StatusOK, 1, nil
}

func (h *Handler) deleteHotel(c *fiber.Ctx) (any, int, int64, error) {
	id := c.Params("id")

	if err := h.hotelService.Delete(c.Context(), id); err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return nil, fiber.StatusNoContent, 0, nil
}

package v1

import (
	"go-booking/internal/dto"
	"go-booking/internal/models"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) listRoom(c *fiber.Ctx) (any, int, int64, error) {
	var filter dto.ListRoomFilter

	if err := c.QueryParser(&filter); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	rooms, count, err := h.roomService.List(c.Context(), filter)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return rooms, fiber.StatusOK, count, nil
}

func (h *Handler) createRoom(c *fiber.Ctx) (any, int, int64, error) {
	var dto dto.CreateRoomRequest
	if err := c.BodyParser(&dto); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	room, err := h.roomService.Create(c.Context(), dto)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return room, fiber.StatusCreated, 1, nil
}

func (h *Handler) updateRoom(c *fiber.Ctx) (any, int, int64, error) {
	id := c.Params("id")

	var room models.Room
	if err := c.BodyParser(&room); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	updatedRoom, err := h.roomService.Update(c.Context(), id, room)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return updatedRoom, fiber.StatusOK, 1, nil
}

func (h *Handler) deleteRoom(c *fiber.Ctx) (any, int, int64, error) {
	id := c.Params("id")

	if err := h.roomService.Delete(c.Context(), id); err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return nil, fiber.StatusNoContent, 0, nil
}

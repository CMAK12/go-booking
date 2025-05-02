package v1

import (
	"go-booking/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) listUser(c *fiber.Ctx) (any, int, int64, error) {
	var filter dto.ListUserFilter

	if err := c.QueryParser(&filter); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	users, count, err := h.userService.List(c.Context(), filter)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return users, fiber.StatusOK, count, nil
}

func (h *Handler) createUser(c *fiber.Ctx) (any, int, int64, error) {
	var req dto.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	user, err := h.userService.Create(c.Context(), req)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return user, fiber.StatusCreated, 1, nil
}

func (h *Handler) updateUser(c *fiber.Ctx) (any, int, int64, error) {
	id := c.Params("id")

	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	user, err := h.userService.Update(c.Context(), id, req)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return user, fiber.StatusOK, 1, nil
}

func (h *Handler) deleteUser(c *fiber.Ctx) (any, int, int64, error) {
	id := c.Params("id")

	if err := h.userService.Delete(c.Context(), id); err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return nil, fiber.StatusNoContent, 0, nil
}

package v1

import (
	"go-booking/internal/dto"
	"go-booking/internal/models"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) listExtraService(c *fiber.Ctx) (any, int, int64, error) {
	var filter dto.ListExtraServiceFilter

	if err := c.QueryParser(&filter); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	extraServices, count, err := h.extraService.List(c.Context(), filter)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return extraServices, fiber.StatusOK, count, nil
}

func (h *Handler) createExtraService(c *fiber.Ctx) (any, int, int64, error) {
	var dto dto.CreateExtraServiceRequest
	if err := c.BodyParser(&dto); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	extraService, err := h.extraService.Create(c.Context(), dto)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return extraService, fiber.StatusCreated, 1, nil
}

func (h *Handler) updateExtraService(c *fiber.Ctx) (any, int, int64, error) {
	id := c.Params("id")

	var extraService models.ExtraService
	if err := c.BodyParser(&extraService); err != nil {
		return nil, fiber.StatusBadRequest, 0, err
	}

	updatedExtraService, err := h.extraService.Update(c.Context(), id, extraService)
	if err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return updatedExtraService, fiber.StatusOK, 1, nil
}

func (h *Handler) deleteExtraService(c *fiber.Ctx) (any, int, int64, error) {
	id := c.Params("id")

	if err := h.extraService.Delete(c.Context(), id); err != nil {
		return nil, fiber.StatusInternalServerError, 0, err
	}

	return nil, fiber.StatusNoContent, 0, nil
}

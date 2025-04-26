package v1

import (
	"encoding/json"
	"net/http"

	"go-booking/internal/dto"
	"go-booking/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
)

func (h *Handler) listExtraService(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	var filter dto.ListExtraServiceFilter

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	if err := decoder.Decode(&filter, r.URL.Query()); err != nil {
		return nil, http.StatusBadRequest, 0, err
	}

	rooms, count, err := h.extraService.List(r.Context(), filter)
	if err != nil {
		return nil, http.StatusInternalServerError, 0, err
	}

	return rooms, http.StatusOK, count, nil
}

func (h *Handler) createExtraService(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	var dto dto.CreateExtraServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return nil, http.StatusBadRequest, 0, err
	}

	extraService, err := h.extraService.Create(r.Context(), dto)
	if err != nil {
		return nil, http.StatusInternalServerError, 0, err
	}

	return extraService, http.StatusCreated, 1, nil
}

func (h *Handler) updateExtraService(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	id := chi.URLParam(r, "id")

	var extraService models.ExtraService
	if err := json.NewDecoder(r.Body).Decode(&extraService); err != nil {
		return nil, http.StatusBadRequest, 0, err
	}

	updatedExtraService, err := h.extraService.Update(r.Context(), id, extraService)
	if err != nil {
		return nil, http.StatusInternalServerError, 0, err
	}

	return updatedExtraService, http.StatusOK, 1, nil
}

func (h *Handler) deleteExtraService(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	id := chi.URLParam(r, "id")

	if err := h.extraService.Delete(r.Context(), id); err != nil {
		return nil, http.StatusInternalServerError, 0, err
	}

	return nil, http.StatusNoContent, 0, nil
}

package v1

import (
	"encoding/json"
	"net/http"

	"go-booking/internal/dto"
	"go-booking/internal/filter"
	"go-booking/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
)

func (h *Handler) listExtraService(w http.ResponseWriter, r *http.Request) {
	var filter filter.ListExtraServiceFilter

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	if err := decoder.Decode(&filter, r.URL.Query()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rooms, count, err := h.extraService.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, rooms, count)
}

func (h *Handler) createExtraService(w http.ResponseWriter, r *http.Request) {
	var dto dto.CreateExtraServiceRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	extraService, err := h.extraService.Create(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, extraService)
}

func (h *Handler) updateExtraService(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var extraService models.ExtraService
	if err := json.NewDecoder(r.Body).Decode(&extraService); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedExtraService, err := h.extraService.Update(r.Context(), id, extraService)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, updatedExtraService)
}

func (h *Handler) deleteExtraService(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.extraService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Extra service deleted successfully"))
}

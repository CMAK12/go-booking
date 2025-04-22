package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) listHotel(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	filter := storage.ListHotelFilter{
		ID:          r.URL.Query().Get("id"),
		Name:        r.URL.Query().Get("name"),
		City:        r.URL.Query().Get("city"),
		Address:     r.URL.Query().Get("address"),
		Description: r.URL.Query().Get("description"),
	}

	if rating := r.URL.Query().Get("rating"); rating != "" {
		ratingFloat, err := strconv.ParseFloat(rating, 64)
		if err != nil {
			http.Error(w, "invalid rating", http.StatusBadRequest)
			return nil, http.StatusBadRequest, 0, err
		}
		filter.Rating = ratingFloat
	}

	hotels, count, err := h.hotelService.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, http.StatusInternalServerError, 0, err
	}

	return hotels, http.StatusOK, count, nil
}

func (h *Handler) createHotel(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	dto := dto.CreateHotelRequest{}
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, http.StatusBadRequest, 0, err
	}

	hotel, err := h.hotelService.Create(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, http.StatusInternalServerError, 0, err
	}

	return hotel, http.StatusCreated, 1, nil
}

func (h *Handler) updateHotel(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	id := chi.URLParam(r, "id")

	hotel := models.Hotel{}
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, http.StatusBadRequest, 0, err
	}

	hotel, err := h.hotelService.Update(r.Context(), id, hotel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, http.StatusInternalServerError, 0, err
	}

	return hotel, http.StatusOK, 1, nil
}

func (h *Handler) deleteHotel(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	id := chi.URLParam(r, "id")

	err := h.hotelService.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, http.StatusInternalServerError, 0, err
	}

	return nil, http.StatusNoContent, 0, nil
}

package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"go-booking/internal/dto"
	"go-booking/internal/models"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
)

func (h *Handler) listHotel(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	var filter dto.ListHotelFilter

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	if err := decoder.Decode(&filter, r.URL.Query()); err != nil {
		return nil, http.StatusBadRequest, 0, err
	}

	if rating := r.URL.Query().Get("rating"); rating != "" {
		ratingFloat, err := strconv.ParseFloat(rating, 64)
		if err != nil {
			return nil, http.StatusBadRequest, 0, err
		}
		filter.Rating = ratingFloat
	}

	hotels, count, err := h.hotelService.List(r.Context(), filter)
	if err != nil {
		return nil, http.StatusInternalServerError, 0, err
	}

	return hotels, http.StatusOK, count, nil
}

func (h *Handler) createHotel(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	dto := dto.CreateHotelRequest{}
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return nil, http.StatusBadRequest, 0, err
	}

	hotel, err := h.hotelService.Create(r.Context(), dto)
	if err != nil {
		return nil, http.StatusInternalServerError, 0, err
	}

	return hotel, http.StatusCreated, 1, nil
}

func (h *Handler) updateHotel(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	id := chi.URLParam(r, "id")

	hotel := models.Hotel{}
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		return nil, http.StatusBadRequest, 0, err
	}

	hotel, err := h.hotelService.Update(r.Context(), id, hotel)
	if err != nil {
		return nil, http.StatusInternalServerError, 0, err
	}

	return hotel, http.StatusOK, 1, nil
}

func (h *Handler) deleteHotel(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	id := chi.URLParam(r, "id")

	err := h.hotelService.Delete(r.Context(), id)
	if err != nil {
		return nil, http.StatusInternalServerError, 0, err
	}

	return nil, http.StatusNoContent, 0, nil
}

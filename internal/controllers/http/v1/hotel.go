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

func (h *Handler) listHotel(w http.ResponseWriter, r *http.Request) {
	var filter dto.ListHotelFilter

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	if err := decoder.Decode(&filter, r.URL.Query()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if rating := r.URL.Query().Get("rating"); rating != "" {
		ratingFloat, err := strconv.ParseFloat(rating, 64)
		if err != nil {
			http.Error(w, "invalid rating", http.StatusBadRequest)
			return
		}
		filter.Rating = ratingFloat
	}

	hotels, count, err := h.hotelService.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, hotels, count)
}

func (h *Handler) createHotel(w http.ResponseWriter, r *http.Request) {
	dto := dto.CreateHotelRequest{}
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hotel, err := h.hotelService.Create(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, hotel)
}

func (h *Handler) updateHotel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	hotel := models.Hotel{}
	if err := json.NewDecoder(r.Body).Decode(&hotel); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hotel, err := h.hotelService.Update(r.Context(), id, hotel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, hotel)
}

func (h *Handler) deleteHotel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.hotelService.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("hotel deleted successfully"))
}

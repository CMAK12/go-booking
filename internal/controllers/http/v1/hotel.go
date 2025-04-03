package v1

import (
	"encoding/json"
	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) listHotel(w http.ResponseWriter, r *http.Request) {
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
			return
		}
		filter.Rating = ratingFloat
	}

	hotels, err := h.hotelService.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(hotels)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
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

	jsonData, err := json.Marshal(hotel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
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

	jsonData, err := json.Marshal(hotel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
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

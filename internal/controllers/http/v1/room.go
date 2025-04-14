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

func (h *Handler) listRoom(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	filter := storage.ListRoomFilter{
		ID:          query.Get("id"),
		HotelID:     query.Get("hotel_id"),
		Name:        query.Get("name"),
		Description: query.Get("description"),
		Available:   query.Get("available"),
	}

	if price := query.Get("price"); price != "" {
		p, err := strconv.Atoi(price)
		if err != nil {
			http.Error(w, "Invalid price value", http.StatusBadRequest)
			return
		}
		filter.Price = p
	}
	if capacity := query.Get("capacity"); capacity != "" {
		c, err := strconv.Atoi(capacity)
		if err != nil {
			http.Error(w, "Invalid capacity value", http.StatusBadRequest)
			return
		}
		filter.Capacity = c
	}
	if quantity := query.Get("quantity"); quantity != "" {
		q, err := strconv.Atoi(quantity)
		if err != nil {
			http.Error(w, "Invalid quantity value", http.StatusBadRequest)
			return
		}
		filter.Quantity = q
	}

	rooms, count, err := h.roomService.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, rooms, count)
}

func (h *Handler) createRoom(w http.ResponseWriter, r *http.Request) {
	var dto dto.CreateRoomRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	room, err := h.roomService.Create(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusCreated, room)
}

func (h *Handler) updateRoom(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var room models.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	room, err := h.roomService.Update(r.Context(), id, room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, room)
}

func (h *Handler) deleteRoom(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.roomService.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Room deleted successfully"))
}

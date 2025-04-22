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

func (h *Handler) listRoom(w http.ResponseWriter, r *http.Request) {
	var filter filter.ListRoomFilter

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	if err := decoder.Decode(&filter, r.URL.Query()); err != nil {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
		return
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

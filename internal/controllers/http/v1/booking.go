package v1

import (
	"encoding/json"
	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) listBooking(w http.ResponseWriter, r *http.Request) {
	filter := storage.ListBookingFilter{
		ID:        r.URL.Query().Get("id"),
		UserID:    r.URL.Query().Get("user_id"),
		RoomID:    r.URL.Query().Get("room_id"),
		StartDate: r.URL.Query().Get("start_date"),
		EndDate:   r.URL.Query().Get("end_date"),
		Status:    models.BookingStatus(r.URL.Query().Get("status")),
	}

	bookings, err := h.bookingService.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(bookings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *Handler) createBooking(w http.ResponseWriter, r *http.Request) {
	var dto dto.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	booking, err := h.bookingService.Create(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func (h *Handler) updateBooking(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var booking dto.UpdateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBooking, err := h.bookingService.Update(r.Context(), id, booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(updatedBooking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *Handler) deleteBooking(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.bookingService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

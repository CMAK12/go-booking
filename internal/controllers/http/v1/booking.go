package v1

import (
	"encoding/json"
	"net/http"

	"go-booking/internal/dto"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
)

func (h *Handler) listBooking(w http.ResponseWriter, r *http.Request) {
	var filter dto.ListBookingFilter

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	if err := decoder.Decode(&filter, r.URL.Query()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bookings, count, err := h.bookingService.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, bookings, count)
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

	writeJSON(w, http.StatusCreated, booking)
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

	writeJSON(w, http.StatusOK, updatedBooking)
}

func (h *Handler) deleteBooking(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := h.bookingService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Booking deleted successfully"))
}

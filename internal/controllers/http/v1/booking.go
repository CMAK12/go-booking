package v1

import (
	"encoding/json"
	"net/http"

	"go-booking/internal/dto"
	"go-booking/internal/models"
	"go-booking/internal/storage"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) listBooking(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	filter := storage.ListBookingFilter{
		ID:        r.URL.Query().Get("id"),
		UserID:    r.URL.Query().Get("user_id"),
		RoomID:    r.URL.Query().Get("room_id"),
		StartDate: r.URL.Query().Get("start_date"),
		EndDate:   r.URL.Query().Get("end_date"),
		Status:    models.BookingStatus(r.URL.Query().Get("status")),
	}

	bookings, count, err := h.bookingService.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, http.StatusInternalServerError, 0, err
	}

	return bookings, http.StatusOK, count, nil
}

func (h *Handler) createBooking(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	var dto dto.CreateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, http.StatusBadRequest, 0, err
	}

	booking, err := h.bookingService.Create(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, http.StatusInternalServerError, 0, err
	}

	return booking, http.StatusCreated, 1, nil
}

func (h *Handler) updateBooking(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	id := chi.URLParam(r, "id")

	var booking dto.UpdateBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, http.StatusBadRequest, 0, err
	}

	updatedBooking, err := h.bookingService.Update(r.Context(), id, booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, http.StatusInternalServerError, 0, err
	}

	return updatedBooking, http.StatusOK, 1, nil
}

func (h *Handler) deleteBooking(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	id := chi.URLParam(r, "id")

	if err := h.bookingService.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, http.StatusInternalServerError, 0, err
	}

	return nil, http.StatusNoContent, 0, nil
}

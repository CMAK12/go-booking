package v1

import (
	"encoding/json"
	"go-booking/internal/dto"
	"go-booking/internal/storage"
	"net/http"
)

func (h *Handler) listBookingService(w http.ResponseWriter, r *http.Request) {
	filter := storage.ListBookingServiceRelationFilter{
		BookingID:      r.URL.Query().Get("booking_id"),
		ExtraServiceID: r.URL.Query().Get("service_id"),
	}

	relations, err := h.bookingServiceRelationService.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(relations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *Handler) createBookingService(w http.ResponseWriter, r *http.Request) {
	var dto dto.CreateBookingServiceRelationRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	relation, err := h.bookingServiceRelationService.Create(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(relation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func (h *Handler) deleteBookingService(w http.ResponseWriter, r *http.Request) {
	bookingID := r.URL.Query().Get("booking_id")
	serviceID := r.URL.Query().Get("service_id")

	if err := h.bookingServiceRelationService.Delete(r.Context(), bookingID, serviceID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Booking service relation deleted successfully"))
}

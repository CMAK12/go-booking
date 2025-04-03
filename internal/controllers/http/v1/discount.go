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

func (h *Handler) listDiscount(w http.ResponseWriter, r *http.Request) {
	filter := storage.ListDiscountFilter{
		ID:      r.URL.Query().Get("id"),
		HotelID: r.URL.Query().Get("hotel_id"),
	}

	if amount := r.URL.Query().Get("amount"); amount != "" {
		parsedAmount, err := strconv.ParseFloat(amount, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		filter.Amount = parsedAmount
	}
	if active := r.URL.Query().Get("active"); active != "" {
		parsedActive, err := strconv.ParseBool(active)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		filter.Active = parsedActive
	}

	discounts, err := h.discountService.List(r.Context(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(discounts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *Handler) createDiscount(w http.ResponseWriter, r *http.Request) {
	var dto dto.CreateDiscountRequest
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	discount, err := h.discountService.Create(r.Context(), dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(discount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
}

func (h *Handler) updateDiscount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var discount models.Discount
	if err := json.NewDecoder(r.Body).Decode(&discount); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedDiscount, err := h.discountService.Update(r.Context(), id, discount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(updatedDiscount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *Handler) deleteDiscount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := h.discountService.Delete(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

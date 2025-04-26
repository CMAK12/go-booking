package v1

import (
	"encoding/json"
	"net/http"

	"go-booking/internal/dto"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
)

func (h *Handler) listUser(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	var filter dto.ListUserFilter

	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	if err := decoder.Decode(&filter, r.URL.Query()); err != nil {
		return nil, http.StatusBadRequest, 0, err
	}

	users, count, err := h.userService.List(r.Context(), filter)
	if err != nil {
		return nil, http.StatusInternalServerError, 0, err
	}

	return users, http.StatusOK, count, nil
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	var req dto.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, http.StatusBadRequest, 0, err
	}

	user, err := h.userService.Create(r.Context(), req)
	if err != nil {
		return nil, http.StatusInternalServerError, 0, err
	}

	return user, http.StatusCreated, 1, nil
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	id := chi.URLParam(r, "id")

	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, http.StatusBadRequest, 0, err
	}

	user, err := h.userService.Update(r.Context(), id, req)
	if err != nil {
		return nil, http.StatusInternalServerError, 0, err
	}

	return user, http.StatusOK, 1, nil
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) (any, int, int64, error) {
	id := chi.URLParam(r, "id")

	if err := h.userService.Delete(r.Context(), id); err != nil {
		return nil, http.StatusInternalServerError, 0, err
	}

	return nil, http.StatusNoContent, 0, nil
}

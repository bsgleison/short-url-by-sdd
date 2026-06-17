package http

import (
	"encoding/json"
	"net/http"

	"github.com/bsgleison/short-url-by-sdd/internal/application/models"
	"github.com/bsgleison/short-url-by-sdd/internal/application/usecase"
)

type CreateShortURLHandler struct {
	createShortURLUseCase *usecase.CreateShortURLUseCase
}

func NewCreateShortURLHandler(createShortURLUseCase *usecase.CreateShortURLUseCase) *CreateShortURLHandler {
	return &CreateShortURLHandler{
		createShortURLUseCase: createShortURLUseCase,
	}
}

func (h *CreateShortURLHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.CreateShortURLInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request body"})
		return
	}

	result, err := h.createShortURLUseCase.Execute(r.Context(), &input)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	if result.HasError {
		writeJSON(w, http.StatusBadRequest, result)
		return
	}

	writeJSON(w, http.StatusCreated, result.Output)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

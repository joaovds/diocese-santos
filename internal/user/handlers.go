package user

import (
	"encoding/json"
	"net/http"
)

type Handlers struct {
	mux      *http.ServeMux
	usecases Usecases
}

func NewHandlers(mux *http.ServeMux) *Handlers {
	userMux := http.NewServeMux()
	mux.Handle("/users/", http.StripPrefix("/users", userMux))

	usecases := NewUserUsecases()

	return &Handlers{mux: userMux, usecases: usecases}
}

// ----- ... -----

func (h *Handlers) SetupRoutes() {
	h.mux.HandleFunc("GET /get-by-id", h.getByID)
}

// ----- ... -----

func (h *Handlers) getByID(w http.ResponseWriter, r *http.Request) {
	user, err := h.usecases.GetByID(1)
	if err.IsError() {
		http.Error(w, err.Message, err.Status)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

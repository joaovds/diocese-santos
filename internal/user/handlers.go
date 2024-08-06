package user

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	h.mux.HandleFunc("GET /{id}", h.getByID)
}

// ----- ... -----

func (h *Handlers) getByID(w http.ResponseWriter, r *http.Request) {
	userIDstr := r.PathValue("id")
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	user, appErr := h.usecases.GetByID(userID)
	if appErr.IsError() {
		http.Error(w, appErr.Message, appErr.Status)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

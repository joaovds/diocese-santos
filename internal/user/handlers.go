package user

import (
	"net/http"

	"github.com/joaovds/diocese-santos/internal/user/usecases"
)

type Handlers struct {
	mux      *http.ServeMux
	usecases *usecases.UserUsecases
}

func NewHandlers(mux *http.ServeMux) *Handlers {
	userMux := http.NewServeMux()
	mux.Handle("/users/", http.StripPrefix("/users", userMux))
	usecases := usecases.NewUserUsecases()
	return &Handlers{mux: userMux, usecases: usecases}
}

// ----- ... -----

func (h *Handlers) SetupRoutes() {
	h.mux.HandleFunc("GET /", h.signIn)
}

// ----- ... -----

func (h *Handlers) signIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
}

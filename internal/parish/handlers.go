package parish

import (
	"encoding/json"
	"net/http"
)

type Handlers struct {
	mux      *http.ServeMux
	usecases Usecases
}

func NewHandlers(mux *http.ServeMux) *Handlers {
	parishMux := http.NewServeMux()
	mux.Handle("/parishes/", http.StripPrefix("/parishes", parishMux))

	usecases := NewParishUsecases()

	return &Handlers{mux: parishMux, usecases: usecases}
}

// ----- ... -----

func (h *Handlers) SetupRoutes() {
	h.mux.HandleFunc("GET /get-by-city", h.getParishesByCity)
}

// ----- ... -----

func (h *Handlers) getParishesByCity(w http.ResponseWriter, r *http.Request) {
	parishes, err := h.usecases.GetParishesByCity([]int{1})
	if err.IsError() {
		http.Error(w, err.Message, err.Status)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(parishes)
}

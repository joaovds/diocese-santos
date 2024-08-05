package liturgy

import (
	"encoding/json"
	"net/http"
)

type Handlers struct {
	mux      *http.ServeMux
	usecases Usecases
}

func NewHandlers(mux *http.ServeMux) *Handlers {
	liturgyMux := http.NewServeMux()
	mux.Handle("/liturgy/", http.StripPrefix("/liturgy", liturgyMux))

	usecases := NewLiturgyUsecases()

	return &Handlers{mux: liturgyMux, usecases: usecases}
}

// ----- ... -----

func (h *Handlers) SetupRoutes() {
	h.mux.HandleFunc("GET /current-liturgical-info", h.getCurrentLiturgicalInfo)
}

// ----- ... -----

func (h *Handlers) getCurrentLiturgicalInfo(w http.ResponseWriter, r *http.Request) {
	currentLiturgicalInfo, err := h.usecases.GetCurrentLiturgicalInfo()
	if err.IsError() {
		http.Error(w, err.Message, err.Status)
		return
	}

	w.WriteHeader(http.StatusUnprocessableEntity)
	json.NewEncoder(w).Encode(currentLiturgicalInfo)
}

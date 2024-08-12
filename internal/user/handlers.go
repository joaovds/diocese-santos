package user

import (
	"net/http"
	"strconv"

	"github.com/joaovds/diocese-santos/pkg/apperr"
	"github.com/joaovds/diocese-santos/pkg/helpers"
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
		helpers.SendHttpResponse(
			w,
			helpers.NewHttpResponseFromError[any](apperr.NewAppError(&INVALID_USER_ID, &UserErrors).SetStatus(http.StatusBadRequest)),
		)
		return
	}

	user, appErr := h.usecases.GetByID(userID)
	if appErr.IsError() {
		helpers.SendHttpResponse(w, helpers.NewHttpResponseFromError[any](appErr))
		return
	}

	helpers.SendHttpResponse(w, helpers.NewHttpResponseFromData(http.StatusOK, user))
}

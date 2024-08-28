package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joaovds/diocese-santos/internal/user/errs"
	"github.com/joaovds/diocese-santos/internal/user/usecases"
	"github.com/joaovds/diocese-santos/pkg/apperr"
	"github.com/joaovds/diocese-santos/pkg/helpers"
)

type Handlers struct {
	mux      *http.ServeMux
	usecases usecases.UserUsecasesContract
}

func NewHandlers(mux *http.ServeMux) *Handlers {
	userMux := http.NewServeMux()
	mux.Handle("/users/", http.StripPrefix("/users", userMux))
	usecases := usecases.NewUserUsecases()
	return &Handlers{mux: userMux, usecases: usecases}
}

func (h *Handlers) SetupRoutes() {
	h.mux.HandleFunc("POST /", h.signIn)
}

// ----- ... -----

func (h *Handlers) signIn(w http.ResponseWriter, r *http.Request) {
	params := usecases.NewSignInUsecaseParams()
	if err := json.NewDecoder(r.Body).Decode(params); err != nil {
		helpers.SendHttpResponse(w, helpers.NewHttpResponseFromError[any](apperr.NewAppError(&errs.INVALID_PARAMS, &errs.UserErrors).SetStatus(http.StatusBadRequest)))
		return
	}
	result, appErr := h.usecases.SignIn(r.Context(), params)
	if appErr.IsError() {
		fmt.Println("aqui", appErr)
		helpers.SendHttpResponse(w, helpers.NewHttpResponseFromError[any](appErr))
		return
	}
	helpers.SendHttpResponse(w, helpers.NewHttpResponseFromData(http.StatusCreated, result))
}

package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/joaovds/diocese-santos/internal/user/errs"
	"github.com/joaovds/diocese-santos/internal/user/usecases"
	"github.com/joaovds/diocese-santos/pkg/apperr"
	"github.com/joaovds/diocese-santos/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewHandlers(t *testing.T) {
	mainMux := http.NewServeMux()
	handlers := NewHandlers(mainMux)
	assert.NotNil(t, handlers)
	assert.NotNil(t, handlers.mux)
	assert.NotNil(t, handlers.usecases)

	routeExists := false
	handlers.mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		routeExists = true
		w.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	mainMux.ServeHTTP(rr, httptest.NewRequest("POST", "/users/", nil))

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.True(t, routeExists, "Expected route /users to be registered in mainMux")
}

func TestHandlers_signIn(t *testing.T) {
	t.Run("should return 400 status", func(t *testing.T) {
		h := &Handlers{mux: http.NewServeMux()}
		h.SetupRoutes()
		req := httptest.NewRequest("POST", "/", nil)
		rr := httptest.NewRecorder()
		h.mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})

	t.Run("should return Invalid Params error", func(t *testing.T) {
		h := &Handlers{mux: http.NewServeMux()}
		h.SetupRoutes()
		req := httptest.NewRequest("POST", "/", nil)
		rr := httptest.NewRecorder()
		h.mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		var body helpers.HttpResponse[any]
		err := json.NewDecoder(rr.Body).Decode(&body)
		assert.Nil(t, err)
		assert.Equal(t, errs.UserErrors[&errs.INVALID_PARAMS], body.Error)
	})

	t.Run("should return Missing param error", func(t *testing.T) {
		h := &Handlers{mux: http.NewServeMux()}
		h.SetupRoutes()

		bodyRequest := strings.NewReader("")
		req := httptest.NewRequest("POST", "/", bodyRequest)
		rr := httptest.NewRecorder()
		h.mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
		var bodyResponse helpers.HttpResponse[any]
		err := json.NewDecoder(rr.Body).Decode(&bodyResponse)
		assert.Nil(t, err)
		assert.Equal(t, errs.UserErrors[&errs.INVALID_PARAMS], bodyResponse.Error)
	})

	t.Run("should return 201 status and result on valid input", func(t *testing.T) {
		mockUsecases := new(usecases.MockUserUsecases)
		mockUsecases.On("SignIn", mock.Anything, mock.AnythingOfType("*usecases.SignInUsecaseParams")).Return(&usecases.SignInUsecaseResult{}, nil)

		h := &Handlers{mux: http.NewServeMux(), usecases: mockUsecases}
		h.SetupRoutes()

		bodyRequest := strings.NewReader(`{"first_name": "valid", "last_name": "valid", "email": "mail", "password": "valid"}`)
		req := httptest.NewRequest("POST", "/", bodyRequest)
		rr := httptest.NewRecorder()
		h.mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		var bodyResponse helpers.HttpResponse[*usecases.SignInUsecaseResult]
		err := json.NewDecoder(rr.Body).Decode(&bodyResponse)
		assert.Nil(t, err)

		mockUsecases.AssertExpectations(t)
	})

	t.Run("should return 500 status on internal error", func(t *testing.T) {
		mockUsecases := new(usecases.MockUserUsecases)
		mockErrorResult := &apperr.AppError{Message: "Internal error", Status: 500}
		mockUsecases.On("SignIn", mock.Anything, mock.AnythingOfType("*usecases.SignInUsecaseParams")).
			Return(nil, mockErrorResult)

		h := &Handlers{mux: http.NewServeMux(), usecases: mockUsecases}
		h.SetupRoutes()

		bodyRequest := strings.NewReader(`{"first_name": "valid", "last_name": "valid", "email": "mail", "password": "valid"}`)
		req := httptest.NewRequest("POST", "/", bodyRequest)
		rr := httptest.NewRecorder()
		h.mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		var bodyResponse helpers.HttpResponse[any]
		err := json.NewDecoder(rr.Body).Decode(&bodyResponse)
		assert.Nil(t, err)
		assert.Equal(t, "Internal error", bodyResponse.Error)

		mockUsecases.AssertExpectations(t)
	})
}

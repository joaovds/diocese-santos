package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joaovds/diocese-santos/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

func TestNewHandlers(t *testing.T) {
	mainMux := http.NewServeMux()
	handlers := NewHandlers(mainMux)
	assert.NotNil(t, handlers)
	assert.NotNil(t, handlers.mux)
	assert.NotNil(t, handlers.usecases)

	routeExists := false
	handlers.mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		routeExists = true
		w.WriteHeader(http.StatusOK)
	})
	rr := httptest.NewRecorder()
	mainMux.ServeHTTP(rr, httptest.NewRequest("GET", "/users/", nil))
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.True(t, routeExists, "Expected route /users to be registered in mainMux")
}

func TestHandlers_getByID(t *testing.T) {
	t.Run("should return the user data in case of success", func(t *testing.T) {
		mockUsecases := new(MockUsecases)
		var noErr *apperr.AppError = nil
		expectedUser := &User{
			ID:        1,
			FirstName: "Carol",
			LastName:  "Wojtyła",
			Email:     "johnpaul.ii@mail.com",
			Password:  "thebestpope",
		}
		mockUsecases.On("GetByID", 1).Return(expectedUser, noErr)

		h := &Handlers{mux: http.NewServeMux(), usecases: mockUsecases}
		h.SetupRoutes()
		req := httptest.NewRequest("GET", "/get-by-id", nil)
		rr := httptest.NewRecorder()
		h.mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var user *User
		err := json.NewDecoder(rr.Body).Decode(&user)
		assert.NoError(t, err)
		assert.Equal(t, expectedUser, user)
		mockUsecases.AssertExpectations(t)
	})

	t.Run("should return an error in the event of an error", func(t *testing.T) {
		mockUsecases := new(MockUsecases)
		errCode := apperr.ErrorCode("TEST_ERR")
		errRes := &apperr.AppError{ErrorCode: &errCode, Message: "testing err", Status: 500}
		var expectedUser *User = nil
		mockUsecases.On("GetByID", 1).Return(expectedUser, errRes)

		h := &Handlers{mux: http.NewServeMux(), usecases: mockUsecases}
		h.SetupRoutes()
		req := httptest.NewRequest("GET", "/get-by-id", nil)
		rr := httptest.NewRecorder()
		h.mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		response := rr.Body.String()
		assert.Contains(t, response, "testing err")
		mockUsecases.AssertExpectations(t)
	})
}

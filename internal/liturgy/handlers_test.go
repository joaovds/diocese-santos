package liturgy

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

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
	mainMux.ServeHTTP(rr, httptest.NewRequest("GET", "/liturgy/", nil))
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.True(t, routeExists, "Expected route /liturgy to be registered in mainMux")
}

func TestHandlers_getCurrentLiturgicalInfo(t *testing.T) {
	t.Run("should return the liturgy data in case of success", func(t *testing.T) {
		mockUsecases := new(MockUsecases)
		var noErr *apperr.AppError = nil
		expectedLiturgy := NewLiturgy("A", "Quaresma", "2° semana", "Roxo")
		mockUsecases.On("GetCurrentLiturgicalInfo").Return(expectedLiturgy, noErr)

		h := &Handlers{mux: http.NewServeMux(), usecases: mockUsecases}
		h.SetupRoutes()
		req := httptest.NewRequest("GET", "/current-liturgical-info", nil)
		rr := httptest.NewRecorder()
		h.mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var actualLiturgy Liturgy
		err := json.NewDecoder(rr.Body).Decode(&actualLiturgy)
		assert.NoError(t, err)
		assert.Equal(t, expectedLiturgy.LiturgicalYear, actualLiturgy.LiturgicalYear)
		assert.Equal(t, expectedLiturgy.LiturgicalSeason, actualLiturgy.LiturgicalSeason)
		assert.Equal(t, expectedLiturgy.LiturgicalWeek, actualLiturgy.LiturgicalWeek)
		assert.Equal(t, expectedLiturgy.LiturgicalColor, actualLiturgy.LiturgicalColor)
		expectedDate := expectedLiturgy.Date.Truncate(time.Second)
		actualDate := actualLiturgy.Date.Truncate(time.Second)
		assert.True(t, expectedDate.Equal(actualDate), "The dates are not equal")
		mockUsecases.AssertExpectations(t)
	})

	t.Run("should return an error in the event of an error", func(t *testing.T) {
		mockUsecases := new(MockUsecases)
		errCode := apperr.ErrorCode("TEST_ERR")
		errRes := &apperr.AppError{ErrorCode: &errCode, Message: "testing err", Status: 500}
		var expectedLiturgy *Liturgy = nil
		mockUsecases.On("GetCurrentLiturgicalInfo").Return(expectedLiturgy, errRes)

		h := &Handlers{mux: http.NewServeMux(), usecases: mockUsecases}
		h.SetupRoutes()
		req := httptest.NewRequest("GET", "/current-liturgical-info", nil)
		rr := httptest.NewRecorder()
		h.mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		response := rr.Body.String()
		assert.Contains(t, response, "testing err")
		mockUsecases.AssertExpectations(t)
	})
}

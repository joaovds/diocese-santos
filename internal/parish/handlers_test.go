package parish

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
	mainMux.ServeHTTP(rr, httptest.NewRequest("GET", "/parishes/", nil))
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.True(t, routeExists, "Expected route /parishes to be registered in mainMux")
}

func TestHandlers_getParishesByCity(t *testing.T) {
	t.Run("should return the parishes data in case of success", func(t *testing.T) {
		mockUsecases := new(MockUsecases)
		var noErr *apperr.AppError = nil
		expectedParishes := []*Parish{
			{
				Church: &Church{
					ID:       1,
					Name:     "São José Operário",
					ImageURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS39k8-fBnhFWanxAm5PZ2QLQ1SQaut1-lrqQ&s",
					Address: &Address{
						Street:       "Rua Santa Lúcia Filippini, 82",
						Neighborhood: "Caraguava",
						City:         "Peruíbe",
						State:        "SP",
						PostalCode:   "11750-000",
						Latitude:     -24.294_751_693_837_465,
						Longitude:    -47.021_398_195_610_026,
					},
					Contact: &Contact{
						Phone: "(13) 3455-3239",
						Email: "saojoseoperario.peruibe@gmail.com",
						Site:  "http://www.diocesedesantos.com.br/",
					},
				},
			},
		}
		mockUsecases.On("GetParishesByCity", []int{1}).Return(expectedParishes, noErr)

		h := &Handlers{mux: http.NewServeMux(), usecases: mockUsecases}
		h.SetupRoutes()
		req := httptest.NewRequest("GET", "/get-by-city", nil)
		rr := httptest.NewRecorder()
		h.mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		var parishes []*Parish
		err := json.NewDecoder(rr.Body).Decode(&parishes)
		assert.NoError(t, err)
		assert.Equal(t, expectedParishes, parishes)
		mockUsecases.AssertExpectations(t)
	})

	t.Run("should return an error in the event of an error", func(t *testing.T) {
		mockUsecases := new(MockUsecases)
		errCode := apperr.ErrorCode("TEST_ERR")
		errRes := &apperr.AppError{ErrorCode: &errCode, Message: "testing err", Status: 500}
		var expectedParishes []*Parish = nil
		mockUsecases.On("GetParishesByCity", []int{1}).Return(expectedParishes, errRes)

		h := &Handlers{mux: http.NewServeMux(), usecases: mockUsecases}
		h.SetupRoutes()
		req := httptest.NewRequest("GET", "/get-by-city", nil)
		rr := httptest.NewRecorder()
		h.mux.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		response := rr.Body.String()
		assert.Contains(t, response, "testing err")
		mockUsecases.AssertExpectations(t)
	})
}

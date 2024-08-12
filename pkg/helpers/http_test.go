package helpers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joaovds/diocese-santos/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

func TestHttpResponse(t *testing.T) {
	t.Run("should return new http response", func(t *testing.T) {
		resp := NewHttpResponse[any]("code", "err", true, 409, nil)
		assert.NotNil(t, resp)
		assert.IsType(t, &HttpResponse[any]{}, resp)
		assert.Equal(t, "code", resp.ErrorCode)
		assert.Equal(t, "err", resp.Error)
		assert.True(t, resp.IsError)
		assert.Equal(t, 409, resp.StatusCode)
		assert.Nil(t, resp.Data)
	})

	t.Run("should return new http response from data", func(t *testing.T) {
		data := map[string]string{"hello": "from data"}
		resp := NewHttpResponseFromData(200, data)
		assert.NotNil(t, resp)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, data, resp.Data)
		assert.False(t, resp.IsError)
		assert.Empty(t, resp.Error)
		assert.Empty(t, resp.ErrorCode)
	})

	t.Run("should return new http response from error", func(t *testing.T) {
		errCode := apperr.ErrorCode("Err")
		err := &apperr.AppError{ErrorCode: &errCode, Message: "from error", Status: 400}

		resp := NewHttpResponseFromError[any](err)
		assert.NotNil(t, resp)
		assert.Equal(t, 400, resp.StatusCode)
		assert.Nil(t, resp.Data)
		assert.True(t, resp.IsError)
		assert.Equal(t, "from error", resp.Error)
		assert.Equal(t, "Err", resp.ErrorCode)
	})

	t.Run("should return new http response from error with default status code", func(t *testing.T) {
		errCode := apperr.ErrorCode("Err")
		err := &apperr.AppError{ErrorCode: &errCode, Message: "from error"}

		resp := NewHttpResponseFromError[any](err)
		assert.NotNil(t, resp)
		assert.Equal(t, 500, resp.StatusCode)
		assert.Nil(t, resp.Data)
		assert.True(t, resp.IsError)
		assert.Equal(t, "from error", resp.Error)
		assert.Equal(t, "Err", resp.ErrorCode)
	})
}

func TestSendHttpResponse(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		response := NewHttpResponseFromData(201, "created example")
		SendHttpResponse(recorder, response)

		expectedOut, _ := json.Marshal(response)
		assert.JSONEq(t, string(expectedOut), recorder.Body.String())
	})

	t.Run("fail", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		failSerialize := make(chan int)
		response := NewHttpResponseFromData(201, failSerialize)
		SendHttpResponse(recorder, response)
		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
	})
}

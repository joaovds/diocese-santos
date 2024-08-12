package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/joaovds/diocese-santos/pkg/apperr"
)

type HttpResponse[T any] struct {
	ErrorCode  string `json:"error_code,omitempty"`
	Error      string `json:"error,omitempty"`
	IsError    bool   `json:"is_error"`
	StatusCode int    `json:"status_code"`
	Data       T      `json:"data,omitempty"`
}

func NewHttpResponse[T any](errorCode, err string, isError bool, statusCode int, data T) *HttpResponse[T] {
	return &HttpResponse[T]{
		ErrorCode:  errorCode,
		Error:      err,
		IsError:    isError,
		StatusCode: statusCode,
		Data:       data,
	}
}

func NewHttpResponseFromData[T any](statusCode int, data T) *HttpResponse[T] {
	return &HttpResponse[T]{
		IsError:    false,
		StatusCode: statusCode,
		Data:       data,
	}
}

func NewHttpResponseFromError[T any](err *apperr.AppError) *HttpResponse[T] {
	statusCode := 500
	if err.Status > 0 {
		statusCode = err.Status
	}

	return &HttpResponse[T]{
		ErrorCode:  string(*err.ErrorCode),
		Error:      err.Error(),
		IsError:    true,
		StatusCode: statusCode,
	}
}

func SendHttpResponse[T any](writer http.ResponseWriter, response *HttpResponse[T]) {
	resJson, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(response.StatusCode)
	writer.Write(resJson)
}

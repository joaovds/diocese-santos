package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/joaovds/diocese-santos/pkg/apperr"
)

type HttpResponse[T any] struct {
	Data       T      `json:"data,omitempty"`
	ErrorCode  string `json:"error_code,omitempty"`
	Error      string `json:"error,omitempty"`
	StatusCode int    `json:"status_code"`
	IsError    bool   `json:"is_error"`
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
	errorCode := apperr.ErrorCode("UNKNOWN")
	if err.Status > 0 {
		statusCode = err.Status
	}
	if err.ErrorCode != nil {
		errorCode = *err.ErrorCode
	}

	return &HttpResponse[T]{
		ErrorCode:  string(errorCode),
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

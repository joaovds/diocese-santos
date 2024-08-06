package helpers

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/joaovds/diocese-santos/pkg/apperr"
)

type HttpResponse struct {
	ErrorCode  string `json:"error_code,omitempty"`
	Error      string `json:"error,omitempty"`
	IsError    bool   `json:"is_error"`
	StatusCode int    `json:"status_code"`
	Data       any    `json:"data,omitempty"`
}

func NewHttpResponse(errorCode, err string, isError bool, statusCode int, data any) *HttpResponse {
	return &HttpResponse{
		ErrorCode:  errorCode,
		Error:      err,
		IsError:    isError,
		StatusCode: statusCode,
		Data:       data,
	}
}

func NewHttpResponseFromData(statusCode int, data any) *HttpResponse {
	return &HttpResponse{
		IsError:    false,
		StatusCode: statusCode,
		Data:       data,
	}
}

func NewHttpResponseFromError(err *apperr.AppError) *HttpResponse {
	statusCode := 500
	if err.Status > 0 {
		statusCode = err.Status
	}

	return &HttpResponse{
		ErrorCode:  string(*err.ErrorCode),
		Error:      err.Error(),
		IsError:    true,
		StatusCode: statusCode,
	}
}

func SendHttpResponse(writer io.Writer, response *HttpResponse) error {
	resJson, err := json.Marshal(response)
	if err != nil {
		fmt.Fprint(writer, err.Error())
		return err
	}

	fmt.Fprint(writer, string(resJson))
	return nil
}

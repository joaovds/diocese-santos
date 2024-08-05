package liturgy

import "github.com/joaovds/diocese-santos/pkg/apperr"

type Usecases interface {
	GetCurrentLiturgicalInfo() (*GetCurrentLiturgicalInfoResponse, *apperr.AppError)
}

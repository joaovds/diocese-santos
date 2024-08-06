package liturgy

import "github.com/joaovds/diocese-santos/pkg/apperr"

type LiturgyUsecases struct{}

func NewLiturgyUsecases() *LiturgyUsecases {
	return &LiturgyUsecases{}
}

// ----- ... -----

func (l *LiturgyUsecases) GetCurrentLiturgicalInfo() (*Liturgy, *apperr.AppError) {
	return NewLiturgy("A", "Tempo Comum", "3° semana", "Verde"), nil
}

package parish

import "github.com/joaovds/diocese-santos/pkg/apperr"

type Usecases interface {
	GetParishesByCity(citiesIDs []int) ([]*Parish, *apperr.AppError)
}

package parish

import "github.com/joaovds/diocese-santos/pkg/apperr"

type ParishUsecases struct{}

func NewParishUsecases() *ParishUsecases {
	return &ParishUsecases{}
}

// ----- ... -----

func (p *ParishUsecases) GetParishesByCity(citiesIDs []int) ([]*Parish, *apperr.AppError) {
	return []*Parish{
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
	}, nil
}

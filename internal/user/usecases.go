package user

import "github.com/joaovds/diocese-santos/pkg/apperr"

type UserUsecases struct{}

func NewUserUsecases() *UserUsecases {
	return &UserUsecases{}
}

// ----- ... -----

func (u *UserUsecases) GetByID(id int) (*User, *apperr.AppError) {
	return &User{
		ID:        1,
		FirstName: "Carol",
		LastName:  "Wojtyła",
		Email:     "johnpaul.ii@mail.com",
		Password:  "",
	}, nil
}

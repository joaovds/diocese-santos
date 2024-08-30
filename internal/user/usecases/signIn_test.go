package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignInParams(t *testing.T) {
	t.Run("validate()", func(t *testing.T) {
		t.Run("should return success when params is valid", func(t *testing.T) {
			validParams := &SignInUsecaseParams{
				FirstName: "John",
				LastName:  "Paul II",
				Email:     "valid_mail",
				Password:  "12345",
			}
			assert.Nil(t, validParams.validate())
		})

		validParams := SignInUsecaseParams{
			FirstName: "John",
			LastName:  "Paul II",
			Email:     "valid_mail",
			Password:  "12345",
		}

		t.Run("should return missing first_name error", func(t *testing.T) {
			invalid := validParams
			invalid.FirstName = ""
			assert.NotNil(t, invalid.validate())
			assert.Equal(t, "Missing field: first_name", invalid.validate().Error())
		})
		t.Run("should return missing last_name error", func(t *testing.T) {
			invalid := validParams
			invalid.LastName = ""
			assert.NotNil(t, invalid.validate())
			assert.Equal(t, "Missing field: last_name", invalid.validate().Error())
		})
		t.Run("should return missing email error", func(t *testing.T) {
			invalid := validParams
			invalid.Email = ""
			assert.NotNil(t, invalid.validate())
			assert.Equal(t, "Missing field: email", invalid.validate().Error())
		})
		t.Run("should return missing password error", func(t *testing.T) {
			invalid := validParams
			invalid.Password = ""
			assert.NotNil(t, invalid.validate())
			assert.Equal(t, "Missing field: password", invalid.validate().Error())
		})
	})
}

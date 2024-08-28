package errs

import (
	"fmt"

	"github.com/joaovds/diocese-santos/pkg/apperr"
)

var (
	INVALID_PARAMS = apperr.ErrorCode("USR_01")
	MISSING_FIELD  = apperr.ErrorCode("USR_02")
)

var UserErrors map[*apperr.ErrorCode]string = map[*apperr.ErrorCode]string{
	&INVALID_PARAMS: "Invalid params",
}

func NewMissingFieldErr(field string) *apperr.AppError {
	return apperr.NewAppError(&MISSING_FIELD, &UserErrors).SetMessage(fmt.Sprintf("Missing field: %s", field))
}

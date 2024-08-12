package user

import "github.com/joaovds/diocese-santos/pkg/apperr"

var (
	INVALID_USER_ID = apperr.ErrorCode("USR_01")
)

var UserErrors map[*apperr.ErrorCode]string = map[*apperr.ErrorCode]string{
	&INVALID_USER_ID: "Invalid user ID",
}

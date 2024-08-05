package apperr

type ErrorCode string

func (e *ErrorCode) String() string {
	return string(*e)
}

type AppError struct {
	ErrorCode *ErrorCode `json:"error_code"`
	Message   string     `json:"message"`
	Status    int        `json:"status_code"`
}

func (a *AppError) Error() string {
	return a.Message
}

func (a *AppError) SetStatus(status int) *AppError {
	a.Status = status
	return a
}

func (a *AppError) SetMessage(message string) *AppError {
	a.Message = message
	return a
}

func (a *AppError) IsNoError() bool {
	return a == nil
}

func (a *AppError) IsError() bool {
	return a != nil
}

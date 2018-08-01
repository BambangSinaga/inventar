package inventar

import "errors"

var (
	ErrContextNil              = errors.New("Context is nil")
	ErrInvalidUsernamePassword = errors.New("Username and password invalid")
	ErrUserNotAuthorized       = errors.New("User not authorized")
	ErrUsernameHasBeenTaken    = errors.New("Username has been taken")
)

type customError struct {
	Message string
}

func (e *customError) Error() string {
	return e.Message
}

func NewCustomError(message string) error {
	return &customError{
		Message: message,
	}
}

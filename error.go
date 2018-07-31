package inventar

import "errors"

var (
	ErrContextNil                 = errors.New("Context is nil")
	ErrUserSubscriptionNotUpdated = errors.New("User subscription is not updated")
	ErrInvalidSubscriptionObject  = errors.New("Invalid item subscription")
	ErrSubscriptionNotFound       = errors.New("Subscription Object Not Found")
	ErrSubscriptionAlreadyExists  = errors.New("Subscription Object Already Exists")

	ErrUserNotAuthorized = errors.New("User not authorized")

	ErrBadRequest = errors.New("Please check your request param or body")
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

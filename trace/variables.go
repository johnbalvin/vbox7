package trace

import (
	"errors"
)

var (
	ErrEmpty        = errors.New("err Empty restult")
	ErrParameter    = errors.New("err not correct parameters")
	ErrMaxAttempt   = errors.New("err Max attemps")
	ErrStatusCode   = errors.New("err Not a correct status code")
	ErrUnknownError = errors.New("err ErrUnknownError")
	ErrRedirect     = errors.New("err redirect")
)

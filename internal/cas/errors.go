package cas

import "strconv"

type ErrConnectionFailed struct {
	Err error
}

func (e *ErrConnectionFailed) Error() string {
	return "cannot connect to cas server"
}

func (e *ErrConnectionFailed) Unwrap() error {
	return e.Err
}

type ErrInvalidCredentials struct{}

func (e *ErrInvalidCredentials) Error() string {
	return "invalid cas credentials"
}

type ErrThrottled struct{}

func (e *ErrThrottled) Error() string {
	return "throttled by cas server"
}

type ErrUnknown struct {
	StatusCode int
}

func (e *ErrUnknown) Error() string {
	return "unknown cas error (" + strconv.Itoa(e.StatusCode) + ")"
}

type ErrUnexpectedResponse struct {
	Err error
}

func (e *ErrUnexpectedResponse) Error() string {
	return "unexpected response from cas server"
}

func (e *ErrUnexpectedResponse) Unwrap() error {
	return e.Err
}

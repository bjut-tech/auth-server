package jwt

type ErrInvalidToken struct{}

func (e *ErrInvalidToken) Error() string {
	return "invalid token"
}

type ErrUnauthorized struct{}

func (e *ErrUnauthorized) Error() string {
	return "unauthorized"
}

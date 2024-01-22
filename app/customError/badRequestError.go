package customError

type BadRequestError struct {
	s string
}

func NewBadRequestError(msg string) error {
	return &BadRequestError{msg}
}

func (b *BadRequestError) Error() string {
	return b.s
}

var ErrorBadRequest = &BadRequestError{}

package errors

import "errors"

var (
	ErrNotFound     = errors.New("resource not found")
	ErrInvalidInput = errors.New("invalid input")
	ErrForeignKey   = errors.New("foreign key violation")
	ErrInternal     = errors.New("internal error")
)

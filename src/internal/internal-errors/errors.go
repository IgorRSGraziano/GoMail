package internalerrors

import "errors"

var (
	ErrInternal = errors.New("internal server error")
)

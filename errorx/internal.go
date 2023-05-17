package errorx

import (
	// "fmt"
	"net/http"

	. "github.com/d2jvkpn/trace_error"
)

// service unavailable
func ErrUnavailable(err error) *Error {
	msg := "service unavailable"

	return NewError(
		err,
		http.StatusServiceUnavailable,
		"service_unavailable",
		Msg(msg),
		Skip(2),
	)
}

// internal server error
func ErrInternal(err error) *Error {
	msg := "internal server error"

	return NewError(
		err,
		http.StatusInternalServerError,
		"internal_server_error",
		Msg(msg),
		Skip(2),
	)
}

// unexpected
func ErrUnexpected(err error) *Error {
	msg := "internal server error"

	return NewError(
		err,
		http.StatusInternalServerError,
		"unexpected",
		Msg(msg),
		Skip(2),
	)
}

// panic, skips example: 5
func ErrPanic(err error, skips ...int) *Error {
	msg := "internal server error"
	skip := 2
	if len(skips) > 0 {
		skip = skips[0]
	}

	return NewError(
		err,
		http.StatusInternalServerError,
		"panic",
		Msg(msg),
		Skip(skip),
	)
}

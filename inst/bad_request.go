package inst

import (
	"fmt"
	"net/http"

	. "github.com/d2jvkpn/trace_error"
)

// no_route
func NotRoute() *Error {
	msg := "no route"

	return NewError(
		fmt.Errorf(msg),
		http.StatusBadRequest,
		"no_route",
		Msg(msg),
		Skip(2),
	)
}

// invalid_query
func ErrQuery(err error) *Error {
	return NewError(
		err,
		http.StatusBadRequest,
		"invalid_query",
		Msg("invalid query"),
		Skip(2),
	)
}

// unmarshal_failed
func ErrUnmarshal(err error) *Error {
	return NewError(
		err,
		http.StatusBadRequest,
		"unmarshal_failed",
		Msg("unmarshal failed"),
		Skip(2),
	)
}

// invalid_argument
func InvalidArg(msg string) *Error {
	if msg == "" {
		msg = "invalid argument"
	}

	return NewError(
		fmt.Errorf(msg),
		http.StatusBadRequest,
		"invalid_argument",
		Msg(msg),
		Skip(2),
	)
}

// not_found
func NotFound(msg string) *Error {
	if msg == "" {
		msg = "not found"
	}

	return NewError(
		fmt.Errorf(msg),
		http.StatusBadRequest,
		"not_found",
		Msg(msg),
		Skip(2),
	)
}

// login_required
func LoginRequired() *Error {
	msg := "login required"

	return NewError(
		fmt.Errorf(msg),
		http.StatusBadRequest,
		"login_required",
		Msg(msg),
		Skip(2),
	)
}

// invalid_token
func InvalidToken() *Error {
	msg := "invalid token"

	return NewError(
		fmt.Errorf(msg),
		http.StatusBadRequest,
		"invalid_token",
		Msg(msg),
		Skip(2),
	)
}

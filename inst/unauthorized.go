package inst

import (
	"fmt"
	"net/http"

	. "github.com/d2jvkpn/trace_error"
)

// invalid_token
func ErrToken(err error) *Error {
	msg := "invalid token"

	return NewError(
		err,
		http.StatusUnauthorized,
		"invalid_token",
		Msg(msg),
		Skip(2),
	)
}

func Unauthorized(msg string) (out *Error) {
	if msg == "" {
		msg = "unauthorized"
	}

	return NewError(
		fmt.Errorf(msg),
		http.StatusUnauthorized,
		"unauthorized",
		Msg(msg),
		Skip(2),
	)
}

func AccessDenied(msg string) (out *Error) {
	if msg == "" {
		msg = "access denied"
	}

	return NewError(
		fmt.Errorf(msg),
		http.StatusUnauthorized,
		"access_denied",
		Msg(msg),
		Skip(2),
	)
}

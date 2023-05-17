package errorx

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
		NoTrace(),
	)
}

// unmarshal_failed
func ErrUnmarshal(err error) *Error {
	return NewError(
		err,
		http.StatusBadRequest,
		"unmarshal_failed",
		Msg("unmarshal failed"),
		NoTrace(),
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

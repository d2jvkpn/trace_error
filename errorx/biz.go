package errorx

import (
	"fmt"
	"net/http"

	. "github.com/d2jvkpn/trace_error"
)

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

// resource_conflict
func ResourceConflict(msg string) *Error {
	if msg == "" {
		msg = "resource conflict"
	}

	return NewError(
		fmt.Errorf(msg),
		http.StatusConflict,
		"resource_conflict",
		Msg(msg),
		Skip(2),
	)
}

// status_conflict
func StatusConflict(msg string) *Error {
	if msg == "" {
		msg = "status conflict"
	}

	return NewError(
		fmt.Errorf(msg),
		http.StatusConflict,
		"status_conflict",
		Msg(msg),
		Skip(2),
	)
}

// remote_call
func ErrRemoteCall(e error, skips ...int) *Error {
	msg := "remote call error"
	skip := 2
	if len(skips) > 0 {
		skip = skips[0]
	}

	return NewError(
		e,
		http.StatusServiceUnavailable,
		"remote_call",
		Msg(msg),
		Skip(skip),
	)
}

// exceeds_limit
func ExceedsLimit(msg string) *Error {
	if msg == "" {
		msg = "exceeds limit"
	}

	return NewError(
		fmt.Errorf(msg),
		http.StatusTooManyRequests,
		"exceeds_limit",
		Msg(msg),
		Skip(2),
	)
}

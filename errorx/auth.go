package errorx

import (
	"fmt"
	"net/http"

	. "github.com/d2jvkpn/trace_error"
)

// login_required
func LoginRequired() *Error {
	msg := "login required"

	return NewError(
		fmt.Errorf(msg),
		http.StatusUnauthorized,
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
		http.StatusUnauthorized,
		"invalid_token",
		Msg(msg),
		Skip(2),
	)
}

// incorrect_api_key
func IncorrectApiKey() *Error {
	msg := "incorrect api key"

	return NewError(
		fmt.Errorf(msg),
		http.StatusUnauthorized,
		"incorrect_api_key",
		Msg(msg),
		Skip(2),
	)
}

// key_is_disabled
func KeyIsDisabled() *Error {
	msg := "key is disabled"

	return NewError(
		fmt.Errorf(msg),
		http.StatusUnauthorized,
		"key_is_disabled",
		Msg(msg),
		Skip(2),
	)
}

// key_has_expired
func KeyHasExpired() *Error {
	msg := "key has expired"

	return NewError(
		fmt.Errorf(msg),
		http.StatusUnauthorized,
		"key_has_expired",
		Msg(msg),
		Skip(2),
	)
}

// auth_failed
func ErrAuthFailed(err error) *Error {
	return NewError(
		err,
		http.StatusUnauthorized,
		"auth_failed",
		Msg("auth failed"),
		Skip(2),
	)
}

// service_access_denied
func ServiceAccessDenied() *Error {
	msg := "service access denied"

	return NewError(
		fmt.Errorf(msg),
		http.StatusUnauthorized,
		"service_access_denied",
		Msg(msg),
		Skip(2),
	)
}

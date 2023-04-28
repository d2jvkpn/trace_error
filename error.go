package this_error

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type Error struct {
	Cause error `json:"cause"`

	CodeInt int    `json:"codeInt"`
	CodeStr string `json:"codeStr"`
	Msg     string `json:"msg"`

	Skip int    `json:"skip"`
	Fn   string `json:"fn"`
	File string `json:"file"`
	Line int    `json:"line"`
}

// type ErrorOption func(*Error) bool
type ErrorOption func(*Error)

func Msg(msg string) ErrorOption {
	return func(e *Error) {
		e.Msg = msg
	}
}

func Skip(skip int) ErrorOption {
	return func(e *Error) {
		e.Skip = skip
	}
}

func NewError(cause error, codeInt int, codeStr string, opts ...ErrorOption) (err *Error) {
	if cause == nil {
		return nil
	}

	err = &Error{Cause: cause, CodeInt: codeInt, CodeStr: codeStr, Msg: "...", Skip: 1}
	for _, v := range opts {
		v(err)
	}

	if err.Skip < 0 {
		return err
	}

	fn, file, line, ok := runtime.Caller(err.Skip)
	if !ok {
		return err
	}

	err.Line = line
	err.Fn = runtime.FuncForPC(fn).Name()
	err.File = filepath.Base(file)

	return err
}

func (err *Error) Update() bool {
	if err.Skip <= 0 {
		return false
	}
	err.Skip -= 1

	fn, file, line, ok := runtime.Caller(err.Skip)
	if !ok {
		return false
	}
	err.Line = line
	err.Fn = runtime.FuncForPC(fn).Name()
	err.File = filepath.Base(file)

	return true
}

/*
func (err *Error) Error() string {
	return fmt.Sprintf(
		"cause: %q, code_int: %d, code_str: %q, msg: %q",
		err.Cause.Error(), err.CodeInt, err.CodeStr, err.Msg,
	)
}
*/

func (err *Error) XMsg(msg string) *Error {
	err.Msg = msg
	return err
}

func (err *Error) XCode(codeInt int, codeStr string) *Error {
	err.CodeInt, err.CodeStr = codeInt, codeStr
	return err
}

func (err *Error) String() string {
	return fmt.Sprintf(
		"cause=%q, code_int=%d, code_str=%q, msg=%q",
		err.Cause.Error(), err.CodeInt, err.CodeStr, err.Msg,
	)
}

func (err *Error) Trace() string {
	if err.Fn == "" {
		return ""
	}

	return fmt.Sprintf(
		"skip=%d, fn=%q, file=%q, line=%d",
		err.Skip, err.Fn, err.File, err.Line,
	)
}

func (err *Error) Describe() string {
	return fmt.Sprintf("Error: %s\nTrace: %s", err, err.Trace())
}

func (err *Error) IntoResponse() *Response {
	return &Response{
		Code: err.CodeStr,
		Msg:  err.Msg,
		Data: map[string]any{},
	}
}

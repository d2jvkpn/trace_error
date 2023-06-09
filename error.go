package trace_error

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Don't alter field CodeStr for predefined Error
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

func NoTrace() ErrorOption {
	return func(e *Error) {
		e.Skip = -1
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

	if err.Skip <= 0 {
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

func (err *Error) Retrace() *Error {
	err.Skip = 1

	fn, file, line, ok := runtime.Caller(err.Skip)
	if !ok {
		return err
	}
	err.Line = line
	err.Fn = runtime.FuncForPC(fn).Name()
	err.File = filepath.Base(file)

	return err
}

/*
func (err *Error) Error() string {
	return fmt.Sprintf(
		"cause: %q, code_int: %d, code_str: %q, msg: %q",
		err.Cause.Error(), err.CodeInt, err.CodeStr, err.Msg,
	)
}
*/

func (err *Error) XCause(e error) *Error {
	if e == nil {
		return err
	}
	err.Cause = e
	return err
}

func (err *Error) XMsg(msg string) *Error {
	err.Msg = msg
	return err
}

func (err *Error) XCode(codeInt int) *Error {
	err.CodeInt = codeInt
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
		"fn=%q, file=%q, line=%d, skip=%d",
		err.Fn, err.File, err.Line, err.Skip,
	)
}

func (err *Error) Describe() string {
	str := err.String()
	trace := err.Trace()

	if trace == "" {
		return str
	}
	return fmt.Sprintf("%s; %s", str, trace)
}

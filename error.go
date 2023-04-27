package this_error

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/google/uuid"
)

type Error struct {
	Cause error `json:"cause"`

	RequestId string `json:"requestId"`
	CodeInt   int    `json:"codeInt"`
	CodeStr   string `json:"codeStr"`
	Msg       string `json:"msg"`

	Skip int    `json:"skip"`
	Fn   string `json:"fn"`
	File string `json:"file"`
	Line int    `json:"line"`
}

// type Option func(*Error) bool
type Option func(*Error)

func Msg(msg string) Option {
	return func(e *Error) {
		e.Msg = msg
	}
}

func Skip(skip int) Option {
	return func(e *Error) {
		e.Skip = skip
	}
}

func Uuid(requestId string) Option {
	return func(e *Error) {
		e.RequestId = requestId
	}
}

func NewError(cause error, codeInt int, codeStr string, opts ...Option) (err *Error) {
	if cause == nil {
		return nil
	}

	err = &Error{Cause: cause, CodeInt: codeInt, CodeStr: codeStr, Msg: "...", Skip: 1}
	for _, v := range opts {
		v(err)
	}

	if err.RequestId == "" {
		if id, e := uuid.NewUUID(); e == nil {
			err.RequestId = id.String()
		}
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

func (err *Error) XRequestId(requestId string) *Error {
	err.RequestId = requestId
	return err
}

func (err *Error) String() string {
	return fmt.Sprintf(
		"cause=%q, request_id=%q, code_int=%d, code_str=%q, msg=%q",
		err.Cause.Error(), err.RequestId, err.CodeInt, err.CodeStr, err.Msg,
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

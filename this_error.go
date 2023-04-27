package this_error

import (
	"fmt"
	"path/filepath"
	"runtime"
)

type Error struct {
	Cause   error  `json:"cause"`
	CodeInt int    `json:"codeInt"`
	CodeStr string `json:"codeStr"`
	Msg     string `json:"msg"`

	Skip int    `json:"skip"`
	Fn   string `json:"fn"`
	File string `json:"file"`
	Line int    `json:"line"`
}

type Option func(*Error) bool

func Msg(msg string) Option {
	return func(e *Error) bool {
		e.Msg = msg
		return true
	}
}

func Skip(skip int) Option {
	return func(e *Error) bool {
		if skip >= 0 {
			e.Skip = skip
			return true
		}
		return false
	}
}

func NewError(cause error, codeInt int, codeStr string, opts ...Option) (err *Error) {
	if cause == nil {
		return nil
	}

	err = &Error{Cause: cause, CodeInt: codeInt, CodeStr: codeStr, Msg: "", Skip: 1}
	for _, v := range opts {
		_ = v(err)
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

func (err *Error) UpdateTrace() bool {
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

func (err *Error) String() string {
	return fmt.Sprintf("CodeInt: %d, CodeStr: %q, Msg: %q", err.CodeInt, err.CodeStr, err.Msg)
}

func (err *Error) Trace() string {
	return fmt.Sprintf(
		"Skip: %d, Fn: %s, File: %q, Line: %d",
		err.CodeInt, err.Fn, err.File, err.Line,
	)
}

func (err *Error) HttpResponse() (int, map[string]any) {
	body := map[string]any{
		"code": err.CodeStr,
		"msg":  err.Msg,
		"data": map[string]any{},
	}

	return err.CodeInt, body
}

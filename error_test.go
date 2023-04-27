package this_error

import (
	"fmt"
	"testing"
)

func fna() (err *Error) {
	e := fmt.Errorf("no user")
	return NewError(e, 400, "not_found")
}

func fnb() (err *Error) {
	e := fmt.Errorf("no user")
	return NewError(e, 400, "not_found", Skip(2))
}

func fnc1() (err *Error) {
	e := fmt.Errorf("no user")
	return NewError(e, 400, "not_found", Skip(2))
}

func fnc2() (err *Error) {
	return fnc1()
}

func fnc3() (err *Error) {
	err = fnc1()
	err.UpdateTrace()
	return err
}

func Test01(t *testing.T) {
	var err *Error

	err = fna()
	fmt.Printf("==> fna\nError: %s\nTrace: %s\n", err, err.Trace())

	err = fnb()
	fmt.Printf("==> fnb\nError: %s\nTrace: %s\n", err, err.Trace())

	err = fnc2()
	fmt.Printf("==> fnc2\nError: %s\nTrace: %s\n", err, err.Trace())

	err = fnc3()
	fmt.Printf("==> fna3\nError: %s\nTrace: %s\n", err, err.Trace())
}

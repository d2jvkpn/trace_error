package this_error

import (
	"fmt"
	"testing"
)

func fna() (err *Error) {
	e := fmt.Errorf("invalid page number")
	return NewError(e, 400, "bad_request")
}

func fnb() (err *Error) {
	e := fmt.Errorf("unmarshal failed")
	return NewError(e, 400, "unmarshal_failed", Skip(2))
}

func fnc1() (err *Error) {
	e := fmt.Errorf("no user")
	return NewError(e, 404, "not_found", Skip(2))
}

func fnc2() (err *Error) {
	return fnc1()
}

func fnc3() (err *Error) {
	err = fnc1()
	err.Update()
	return err
}

func func4() (err *Error) {
	e := fmt.Errorf("an error")
	return NewError(e, 503, "service_unavailable", Skip(-1))
}

func Test01(t *testing.T) {
	var err *Error

	err = fna()
	fmt.Printf("==> fna\n%s\n", err.Describe())

	err = fnb()
	fmt.Printf("==> fnb\n%s\n", err.Describe())

	err = fnc2()
	fmt.Printf("==> fnc2\n%s\n", err.Describe())

	err = fnc3()
	fmt.Printf("==> fna3\n%s\n", err.Describe())

	err = func4()
	fmt.Printf("==> fna4\n%s\n", err.Describe())
}

func Test02(t *testing.T) {
	check := func(d any) {
		fmt.Println(d == nil)
	}

	check(nil)
}

func Test03(t *testing.T) {
	e := fmt.Errorf("an error")
	err := NewError(e, 503, "service_unavailable")
	fmt.Println(err.Describe())

	err = NewError(e, 503, "service_unavailable", Skip(-4))
	fmt.Println(err.Describe())
}

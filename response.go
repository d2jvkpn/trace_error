package this_error

import (
	"encoding/json"
)

type Response struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (res Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(res)
}

func OkResponse(data any) Response {
	return Response{
		Code: "ok",
		Msg:  "ok",
		Data: data,
	}
}

func FromError(err *Error) Response {
	return Response{
		Code: err.CodeStr,
		Msg:  err.Msg,
		Data: map[string]any{},
	}
}

func (err *Error) HttpJSON() (int, json.Marshaler) {
	return err.CodeInt, FromError(err)
}

func (res *Response) HttpJSON() (int, json.Marshaler) {
	return 200, res
}

package this_error

import (
	"encoding/json"
	"net/http"
)

type HttpJSON interface {
	HttpJSON() (int, json.Marshaler)
}

func FromError(err *Error) Response {
	return Response{
		Code: err.CodeStr,
		Msg:  err.Msg,
		Data: map[string]any{},
	}
}

func (err *Error) IntoResponse() Response {
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
	if res.Data == nil {
		res.Data = map[string]any{}
	}

	return http.StatusOK, res
}

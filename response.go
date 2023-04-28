package this_error

import (
	// "encoding/json"

	"github.com/google/uuid"
)

type Response struct {
	RequestId string `json:"requestId"`
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	Data      any    `json:"data"`
}

type ResponeOption func(*Response)

/*
func (res *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(res)
}
*/

func (res *Response) XRequestId(requestId string) *Response {
	res.RequestId = requestId
	return res
}

func RequestId(requestId string) ResponeOption {
	return func(res *Response) {
		res.RequestId = requestId
	}
}

func NewResponse(data any, opts ...ResponeOption) Response {
	res := Response{Code: "ok", Msg: "ok", Data: data}

	for _, opt := range opts {
		opt(&res)
	}

	if res.Data == nil {
		res.Data = map[string]any{}
	}

	if res.RequestId == "" {
		if id, e := uuid.NewUUID(); e == nil {
			res.RequestId = id.String()
		}
	}

	return res
}

func (err *Error) IntoResponse(opts ...ResponeOption) Response {
	res := Response{
		Code: err.CodeStr,
		Msg:  err.Msg,
		Data: map[string]any{},
	}

	for _, opt := range opts {
		opt(&res)
	}

	return res
}

func FromError(err *Error, opts ...ResponeOption) Response {
	res := Response{
		Code: err.CodeStr,
		Msg:  err.Msg,
		Data: map[string]any{},
	}

	for _, opt := range opts {
		opt(&res)
	}

	return res
}

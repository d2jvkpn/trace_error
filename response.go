package this_error

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Response struct {
	RequestId string `json:"requestId"`
	Code      string `json:"code"`
	Msg       string `json:"msg"`
	Data      any    `json:"data"`
}

func (res *Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(res)
}

func NewResponse(data any) Response {
	res := Response{Code: "ok", Msg: "ok", Data: data}

	if res.Data == nil {
		res.Data = map[string]any{}
	}

	if id, e := uuid.NewUUID(); e == nil {
		res.RequestId = id.String()
	}

	return res
}

func FromError(err *Error) Response {
	return Response{
		RequestId: err.RequestId,
		Code:      err.CodeStr,
		Msg:       err.Msg,
		Data:      map[string]any{},
	}
}

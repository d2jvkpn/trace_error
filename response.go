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

func NewResponse(data any) Response {
	return Response{
		Code: "ok",
		Msg:  "ok",
		Data: data,
	}
}

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

func (res Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(res)
}

func NewResponse(data any) Response {
	res := Response{
		Code: "ok",
		Msg:  "ok",
		Data: data,
	}

	if id, e := uuid.NewUUID(); e == nil {
		res.RequestId = id.String()
	}

	return res
}

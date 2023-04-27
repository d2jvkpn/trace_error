package this_error

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestResponse(t *testing.T) {
	res := NewResponse(nil)

	bts, _ := json.Marshal(res)
	fmt.Println(string(bts))
}

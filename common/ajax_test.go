package common

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Obj struct {
	Str string `json:"str"`
	Num int    `json:"num"`
}

func TestJsonResp(t *testing.T) {
	arr := [2]*Obj{&Obj{"str1", 1}, &Obj{"str2", 2}}
	b, err := json.Marshal(&Resp{
		Data:arr,
	})
	if err != nil {
		t.Error("encoding faild")
	}
	fmt.Println(string(b))
	var Result Resp
	json.Unmarshal(b, &Result)

}

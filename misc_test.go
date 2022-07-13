package main

import (
	"encoding/json"
	"fmt"
	"github.com/hexcraft-biz/misc/xtime"
	"testing"
)

func TestXtime(t *testing.T) {
	type TimeTesting struct {
		Case1 xtime.NullTimeRFC3339 `json:"case1"`
		Case2 xtime.NullTimeRFC3339 `json:"case2"`
		Case3 xtime.NullTimeRFC3339 `json:"case3"`
	}

	var tt TimeTesting
	jsonStr := []byte(`{
		"case1": "2022-06-16T04:00:00Z",
		"case2": "",
		"case3": null
	}`)

	if err := json.Unmarshal(jsonStr, &tt); err != nil {
		fmt.Println(err.Error())
	} else if js, err := json.Marshal(tt); err != nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(tt.Case1)
		fmt.Println(tt.Case2)
		fmt.Println(tt.Case3)
		fmt.Println(string(js))
	}
}

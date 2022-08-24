package main

import (
	"encoding/json"
	"fmt"
	"github.com/hexcraft-biz/misc/xtime"
	"github.com/hexcraft-biz/misc/xuuid"
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

func TestXuuid(t *testing.T) {
	type XuuidTesting struct {
		Case1 xuuid.UUID  `json:"case1"`
		Case2 xuuid.UUID  `json:"case2"`
		Case3 xuuid.UUID  `json:"case3"`
		Case4 xuuid.UUID  `json:"case4"`
		Case5 *xuuid.UUID `json:"case5"`
		Case6 *xuuid.UUID `json:"case6"`
		Case7 *xuuid.UUID `json:"case7"`
		Case8 *xuuid.UUID `json:"case8"`
	}

	var xt XuuidTesting
	jsonStr := []byte(`{
		"case1": "",
		"case2": null,
		"case5": "",
		"case6": null
	}`)
	if err := json.Unmarshal(jsonStr, &xt); err != nil {
		fmt.Println(err.Error())
	} else if js, err := json.Marshal(xt); err != nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(xt.Case1)
		fmt.Println(xt.Case2)
		fmt.Println(xt.Case3)
		fmt.Println(string(js))
	}
}

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
		Case1 *xtime.TimeStartedRFC3339 `json:"case1"`
		Case2 *xtime.TimeStartedRFC3339 `json:"case2"`
		Case3 *xtime.TimeExpiredRFC3339 `json:"case3"`
		Case4 *xtime.TimeExpiredRFC3339 `json:"case4"`
		Case5 xtime.TimeStartedRFC3339  `json:"case5"`
		Case6 xtime.TimeStartedRFC3339  `json:"case6"`
		Case7 xtime.TimeExpiredRFC3339  `json:"case7"`
		Case8 xtime.TimeExpiredRFC3339  `json:"case8"`
	}

	var tt TimeTesting
	jsonStr := []byte(`{
		"case1": "",
		"case2": null,
		"case3": "",
		"case4": null,
		"case5": "",
		"case6": null,
		"case7": "",
		"case8": null
	}`)

	if err := json.Unmarshal(jsonStr, &tt); err != nil {
		fmt.Println(err.Error())
	} else if js, err := json.Marshal(tt); err != nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(tt.Case1)
		fmt.Println(tt.Case2)
		fmt.Println(tt.Case3)
		fmt.Println(tt.Case4)
		fmt.Println(tt.Case5)
		fmt.Println(tt.Case6)
		fmt.Println(tt.Case7)
		fmt.Println(tt.Case8)
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

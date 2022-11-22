package main

import (
	"encoding/json"
	"fmt"
	"github.com/hexcraft-biz/misc/xtime"
	"github.com/hexcraft-biz/misc/xuuid"
	"testing"
	"time"
)

func TestMysqlTime(t *testing.T) {
	type TimeTest struct {
		Case1 xtime.Time `json:"time_missing"`
		Case2 xtime.Time `json:"time_null"`
		Case3 xtime.Time `json:"time_common"`
	}

	var tt TimeTest
	jsonStr := []byte(`{
		"time_null": null,
		"time_common": "2023-12-01T00:00:00Z"
	}`)

	if err := json.Unmarshal(jsonStr, &tt); err != nil {
		fmt.Println(err.Error())
	} else if js, err := json.MarshalIndent(tt, "", "\t"); err != nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(tt.Case1.Value())
		fmt.Println(tt.Case2.Value())
		fmt.Println(tt.Case3.Value())
		fmt.Println(string(js))
		fmt.Println(tt.Case3.Add(time.Second * 86400))
	}
}

func TestXuuid(t *testing.T) {
	type XuuidTesting struct {
		Case1 xuuid.Wildcard `json:"case1"`
		Case2 xuuid.Wildcard `json:"case2"`
		Case3 xuuid.Wildcard `json:"case3"`
		Case4 xuuid.Wildcard `json:"case4"`
		Case5 xuuid.UUID     `json:"case5"`
	}

	var xt XuuidTesting
	jsonStr := []byte(`{
		"case1": "c973e6dc-c2ea-46ef-b1af-65653d8df62a",
		"case2": "test",
		"case3": "",
		"case4": null,
		"case5": "c973e6dc-c2ea-46ef-b1af-65653d8df62a"
	}`)
	if err := json.Unmarshal(jsonStr, &xt); err != nil {
		fmt.Println(err.Error())
	} else if js, err := json.Marshal(xt); err != nil {
		t.Fatal(err.Error())
	} else {
		fmt.Println(xt.Case1.Value())
		fmt.Println(xt.Case2.Value())
		fmt.Println(xt.Case3.Value())
		fmt.Println(xt.Case4.Value())
		fmt.Println(xt.Case5.Value())
		fmt.Println(string(js))
	}
}

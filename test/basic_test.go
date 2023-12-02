package main

import (
	"fmt"
	"testing"

	"github.com/hexcraft-biz/misc/basic"
)

func TestBasic(t *testing.T) {
	for i := 0; i < 10; i += 1 {
		fmt.Println(basic.GenStringWithCharset(8, basic.DefCharsetNumber|basic.DefCharsetLowercase|basic.DefCharsetUppercase))
	}
}

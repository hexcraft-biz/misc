package main

import (
	"fmt"
	"testing"

	"github.com/hexcraft-biz/misc/imgin"
)

type input struct {
	imgin.Imgin
}

func TestImgin(t *testing.T) {
	i := new(input)
	i.Src = "https://img.freepik.com/premium-vector/runes-magic-symbol-ancient-gothic-sign_86689-135.jpg"
	if err := i.Validate(); err != nil {
		fmt.Println("Test 1 error:")
		fmt.Println(err.O())
	}

	i.Src = "data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"
	if err := i.Validate(); err != nil {
		fmt.Println("Test 2 error:")
		fmt.Println(err.O())
	}

	i.Src = "test.jpg"
	i.DirUploads = "../test/"
	if err := i.Validate(); err != nil {
		fmt.Println("Test 3 error:")
		fmt.Println(err.O())
	}
}

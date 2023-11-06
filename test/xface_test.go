package main

import (
	"fmt"
	"testing"

	"github.com/hexcraft-biz/misc/xface"
)

func TestXface(t *testing.T) {
	f1 := new(xface.Descriptor)
	f2 := new(xface.Descriptor)
	fmt.Println(f1.DistWithFace(f2))
}

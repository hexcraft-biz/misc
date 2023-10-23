package main

import (
	"fmt"
	"testing"

	"github.com/hexcraft-biz/misc/scope"
)

func TestScope(t *testing.T) {
	c1 := ""
	c2 := "user.prototype user.profile user.friends.readonly"

	r1 := []string{"user.prototype"}
	r2 := []string{"user.friends.readonly", "user.friends.write"}

	c1scope := scope.New(c1)
	c2scope := scope.New(c2)
	r1scope := scope.New(r1)
	r2scope := scope.New(r2)

	fmt.Println("c1:", c1scope, "r2:", r2scope, "HasOneOf:", c1scope.HasOneOf(r2scope))
	fmt.Println("c1:", c1scope, "r2:", r2scope, "Contains:", c1scope.Contains(r2scope))
	fmt.Println("c2:", c2scope, "r1:", r1scope, "HasOneOf:", c2scope.HasOneOf(r1scope))
	fmt.Println("c2:", c2scope, "r1:", r1scope, "Contains:", c2scope.Contains(r1scope))
	fmt.Println("c2:", c2scope, "r2:", r2scope, "HasOneOf:", c2scope.HasOneOf(r2scope))
	fmt.Println("c2:", c2scope, "r2:", r2scope, "Contains:", c2scope.Contains(r2scope))
}

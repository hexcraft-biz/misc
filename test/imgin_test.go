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
	if err := input.Validate(); err != nil {
		fmt.Println(err.O())
	}

	i.Src = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAB4AAAAeCAIAAAC0Ujn1AAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAEDSURBVEhLtZJBEoMwDAP7lr6nn+0LqUGChsVOwoGdvTSSNRz6Wh7jxvT7+wn9Y4LZae0e+rXLeBqjh45rBtOYgy4V9KYxlOpqRjmNiY4+uJBP41gOI5BM40w620AknTVwGgfSWQMK0tnOaRpV6ewCatLZxn8aJemsAGXp7JhGLBX1wYlUtE4jkIpnwKGM9xeepG7mwblMpl2/CUbCJ7+6CnQzAw5lvD/8DxGIpbMClKWzdjpASTq7gJp0tnGaDlCVzhpQkM52OB3gQDrbQCSdNSTTAc7kMAL5dIDjjj64UE4HmEh1NaM3HWAIulQwmA4wd+i4ZjwdYDR00GqWsyPrizLD76QCPOHqP2cAAAAAElFTkSuQmCC"
	if err := input.Validate(); err != nil {
		fmt.Println(err.O())
	}

	i.Src = "test.jpg"
	if err := input.Validate(); err != nil {
		fmt.Println(err.O())
	}
}

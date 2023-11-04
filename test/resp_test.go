package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"testing"

	"github.com/hexcraft-biz/misc/resp"
)

func TestResp(t *testing.T) {
	PrintResult(fromRespNew())
	PrintResult(fromRespNewError())
	PrintResult(fromRespNewErrorWithMessage())
	PrintResult(sql.ErrNoRows)
}

type Result struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func fromRespNew() *resp.Resp {
	return resp.New(http.StatusOK, &Result{Id: 1, Name: "HEXCraft"})
}

func fromRespNewError() error {
	return resp.NewError(http.StatusNotFound, sql.ErrNoRows, nil)
}

func fromRespNewErrorWithMessage() error {
	return resp.NewErrorWithMessage(http.StatusInternalServerError, "error message", nil)
}

func PrintResult(err error) {
	r := resp.Assert(err)
	if r != nil {
		fmt.Println("Code:", r.StatusCode, "Message:", r.Payload.Message, "Result:", r.Result)
	} else {
		fmt.Println("r is nil")
	}
}

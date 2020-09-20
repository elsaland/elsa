package main

import (
	"fmt"
	"io/ioutil"

	"github.com/lithdew/quickjs"
)

func DoneNS(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	switch args[0].Int32() {
	case FSRead:
		file := args[1].String()
		dat, e := ioutil.ReadFile(file)
		if e != nil {
			panic(e)
		}
		return ctx.String(string(dat))
	case Log:
		fmt.Println(args[1])
		return ctx.Null()
	default:
		return ctx.Null()
	}
}

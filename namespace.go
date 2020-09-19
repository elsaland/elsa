package main

import (
	"fmt"
	"io/ioutil"

	"github.com/lithdew/quickjs"
)

func DoneNS(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	switch args[0].String() {
	case "readFile":
		file := args[1].String()
		dat, e := ioutil.ReadFile(file)
		if e != nil {
			panic(e)
		}
		return ctx.String(string(dat))
	case "console":
		fmt.Println(args[1])
		return ctx.Null()
	default:
		return ctx.Null()
	}
}

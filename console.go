package main

import (
	"fmt"

	"github.com/lithdew/quickjs"
)

func Console(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	fmt.Println(args[0])
	return ctx.Null()
}

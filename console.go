package main

import (
	"encoding/json"
	"fmt"

	"github.com/lithdew/quickjs"
)

// ConsoleLog console.log bindings to quickjs engine
func ConsoleLog(ctx *quickjs.Context, value []quickjs.Value) quickjs.Value {
	data := value[1]
	var result interface{}
	json.Unmarshal([]byte(data.String()), &result)
	prty, _ := Marshal(result)
	fmt.Println(string(prty))
	return ctx.Null()
}

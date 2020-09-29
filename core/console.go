package core

import (
	"encoding/json"
	"fmt"

	"github.com/elsaland/quickjs"
	"github.com/fatih/color"
)

var f = NewFormatter()

// ConsoleLog console.log bindings to quickjs engine
func ConsoleLog(ctx *quickjs.Context, value []quickjs.Value) quickjs.Value {
	data := value[2]
	dataType := value[1].String()
	var result interface{}
	switch dataType {
	case "string":
		fmt.Println(data.String())
	case "function":
		fmt.Println(color.New(color.FgCyan).SprintFunc()(data.String()))
	case "bigint":
		fmt.Println(color.New(color.FgYellow).SprintFunc()(data.BigInt()))
	case "number":
		fmt.Println(color.New(color.FgYellow).SprintFunc()(data.Int32()))
	default:
		json.Unmarshal([]byte(data.String()), &result)
		prty, _ := f.Marshal(result)
		fmt.Println(string(prty))
	}

	return ctx.Null()
}

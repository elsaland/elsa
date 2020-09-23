package core

import (
	"encoding/json"
	"fmt"

	"github.com/lithdew/quickjs"
)

var f = NewFormatter()

// ConsoleLog console.log bindings to quickjs engine
func ConsoleLog(ctx *quickjs.Context, value []quickjs.Value) quickjs.Value {
	data := value[1]
	var result interface{}
	if data.IsFunction() || data.IsConstructor() {
		// TODO: log with Function name
		result = data.String()
		fmt.Println(result)
	} else {
		json.Unmarshal([]byte(data.String()), &result)

		prty, _ := f.Marshal(result)
		fmt.Println(string(prty))
	}
	return ctx.Null()
}

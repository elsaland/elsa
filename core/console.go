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
    fmt.Fprintln(color.Output, color.New(color.FgCyan).SprintFunc()(data.String()))
  case "bigint":
    fmt.Fprintln(color.Output, color.New(color.FgYellow).SprintFunc()(data.BigInt()))
  case "number":
    fmt.Fprintln(color.Output, color.New(color.FgYellow).SprintFunc()(data.Int32()))
  default:
    json.Unmarshal([]byte(data.String()), &result)
    prty, _ := f.Marshal(result)
    fmt.Fprintln(color.Output, string(prty))
  }

  return ctx.Null()
}

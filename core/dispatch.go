package core

import (
	"io/ioutil"
	"os"

	"github.com/elsaland/elsa/cmd"
	"github.com/lithdew/quickjs"
)

func ElsaNS(perms cmd.Perms) func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	return func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		switch args[0].Int32() {
		case FSRead:
			if !perms.Fs {
				LogError("Perms Error: ", "Filesystem access is blocked.")
				os.Exit(1)
			}
			file := args[1].String()
			dat, e := ioutil.ReadFile(file)
			if e != nil {
				panic(e)
			}
			val := ctx.String(string(dat))
			defer val.Free()
			return val
		case Log:
			return ConsoleLog(ctx, args)
		case Plugin:
			plugin := args[1].String()
			input := args[2].String()
			dat := (OpenPlugin(plugin, input)).(string)
			val := ctx.String(dat)
			defer val.Free()
			return val
		default:
			return ctx.Null()
		}
	}
}

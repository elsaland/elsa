package core

import (
	"os"

	"github.com/elsaland/elsa/core/ops"
	"github.com/spf13/afero"

	"github.com/elsaland/elsa/cmd"
	"github.com/lithdew/quickjs"
)

func ElsaNS(perms cmd.Perms) func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	var fs = ops.FsDriver{
		Fs:    afero.NewOsFs(),
		Perms: perms,
	}
	return func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		switch args[0].Int32() {
		case FSRead:
			CheckFs(perms)
			file := args[1]
			val := fs.ReadFile(ctx, file)
			return val
		case FSExists:
			CheckFs(perms)
			file := args[1]
			val := fs.Exists(ctx, file)
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

func CheckFs(perms cmd.Perms) {
	if !perms.Fs {
		LogError("Perms Error: ", "Filesystem access is blocked.")
		os.Exit(1)
	}
}

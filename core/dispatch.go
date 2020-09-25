package core

import (
	"os"

	"github.com/elsaland/elsa/core/ops"
	"github.com/spf13/afero"
	"github.com/elsaland/elsa/cmd"
  "github.com/elsaland/quickjs"
)

func ElsaSendNS(elsa *Elsa) func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	var fs = ops.FsDriver{
		Fs:    afero.NewOsFs(),
		Perms: elsa.flags,
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
		case FSWrite:
			CheckFs(perms)
			file := args[1]
			contents := args[2]
			val := fs.WriteFile(ctx, file, contents)
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
		case Fetch:
      one := args[1]
      elsa.recv(one, ctx.String("Hello World"))
			return ctx.Null()
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

func ElsaRecvNS(elsa *Elsa) func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	return func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		fn := args[0]
		elsa.recv = func(id quickjs.Value, val quickjs.Value) {
      result := fn.Call(id, val)
      defer result.Free()
    }
		return ctx.Null()
	}
}

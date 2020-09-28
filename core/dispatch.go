package core

import (
	"os"

	"github.com/elsaland/elsa/cmd"
	"github.com/elsaland/elsa/core/ops"
	"github.com/elsaland/quickjs"
	"github.com/spf13/afero"
)

func ElsaSendNS(elsa *Elsa) func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	var fs = ops.FsDriver{
		Fs:    afero.NewOsFs(),
		Perms: elsa.Perms,
	}
	return func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		switch args[0].Int32() {
		case FSRead:
			CheckFs(elsa.Perms)
			file := args[1]
			val := fs.ReadFile(ctx, file)
			return val
		case FSExists:
			CheckFs(elsa.Perms)
			file := args[1]
			val := fs.Exists(ctx, file)
			return val
		case FSWrite:
			CheckFs(elsa.Perms)
			file := args[1]
			contents := args[2]
			val := fs.WriteFile(ctx, file, contents)
			return val
		case FSCwd:
			CheckFs(elsa.Perms)
			val := fs.Cwd(ctx)
			return val
		case FSStats:
			CheckFs(elsa.Perms)
			file := args[1]
			val := fs.Stats(ctx, file)
			return val
		case FSRemove:
			CheckFs(elsa.Perms)
			file := args[1]
			val := fs.Remove(ctx, file)
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
			url := args[2]
			body := ops.Fetch(ctx, url)
			elsa.Recv(one, body)
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
		elsa.Recv = func(id quickjs.Value, val quickjs.Value) {
			result := fn.Call(id, val)
			defer result.Free()
		}
		return ctx.Null()
	}
}

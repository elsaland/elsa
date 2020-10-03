package core

import (
	"fmt"
	"os"

	"github.com/elsaland/elsa/util"

	"github.com/elsaland/elsa/core/ops"
	"github.com/elsaland/elsa/core/options"
	"github.com/elsaland/quickjs"
	"github.com/spf13/afero"
)

// ElsaSendNS Native function corresponding to the Javascript global `__send`
// It is binded with `__send` and accepts arguments including op ID
func ElsaSendNS(elsa *options.Elsa) func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	// Create a new file system driver
	var fs = ops.FsDriver{
		// NOTE: afero can also be used to create in-memory file system
		// it can be a feature to provide in the future
		Fs:    afero.NewOsFs(),
		Perms: elsa.Perms,
	}
	// The returned function handles the op and execute corresponding native code
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
		case FSMkdir:
			CheckFs(elsa.Perms)
			file := args[1]
			val := fs.Mkdir(ctx, file)
			return val
		default:
			return ctx.Null()
		}
	}
}

// CheckFs utility to check whether file system access is avaliable or not
func CheckFs(perms *options.Perms) {
	if !perms.Fs {
		util.LogError("Perms Error: ", "Filesystem access is blocked.")
		os.Exit(1)
	}
}

// ElsaRecvNS Native function corresponding to the Javascript global `__recv`
// It is binded with `__recv` and accepts arguments including recv ID of the async function
func ElsaRecvNS(elsa *options.Elsa) func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	// the returned function handles the __recv behaviour
	// It is capable of calling the callback for a particular async op after it has finished
	return func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		fn := args[0]
		if elsa.Recv != nil {
			ctx.ThrowError(fmt.Errorf("recv cannot be called more than once"))
			return ctx.Null()
		}
		elsa.Recv = func(id quickjs.Value, val quickjs.Value) {
			result := fn.Call(id, val)
			defer result.Free()
		}
		return ctx.Null()
	}
}

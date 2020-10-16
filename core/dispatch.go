package core

import (
	"fmt"
	"os"
	"sync"

	"github.com/elsaland/elsa/util"

	"github.com/elsaland/elsa/core/ops"
	"github.com/elsaland/elsa/core/options"
	"github.com/elsaland/quickjs"
	"github.com/spf13/afero"
)

// ElsaSendNS Native function corresponding to the Javascript global `__send`
// It is binded with `__send` and accepts arguments including op ID
func ElsaSendNS(elsa *options.Elsa, wg *sync.WaitGroup) func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
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
		case FSStat:
			CheckFs(elsa.Perms)
			file := args[1]
			val := fs.Stat(ctx, file)
			return val
		case FSRemove:
			CheckFs(elsa.Perms)
			file := args[1]
			val := fs.Remove(ctx, file)
			return val
		case Log:
			return ConsoleLog(ctx, args)
		case Timers:
			timeout := args[2]
			one := args[1]
			cb := func() {
				obj := ctx.Object()
				defer obj.Free()
				obj.Set("ok", ctx.Bool(true))
				elsa.Recv(one, obj)
			}
			wg.Add(1)
			ops.SetTimeout(ctx, timeout.Int64(), cb, wg)
			return ctx.Null()
		case Plugin:
			plugin := args[1].String()
			input := args[2].String()
			dat := (OpenPlugin(plugin, input)).(string)
			val := ctx.String(dat)
			defer val.Free()
			return val
		case Fetch:
			CheckNet(elsa.Perms)
			one := args[1]
			url := args[2]
			body := ops.Fetch(ctx, url)
			obj := ctx.Object()
			defer obj.Free()
			obj.Set("ok", body)
			elsa.Recv(one, obj)
			return ctx.Null()
		case Serve:
			id := args[1]
			url := args[2]
			cb := func(res quickjs.Value) string {
				obj := ctx.Object()
				defer obj.Free()
				obj.Set("ok", res)
				rtrn := elsa.Recv(id, res)
				return rtrn.String()
			}
			ops.Serve(ctx, cb, id, url)
			return ctx.Null()
		case FSMkdir:
			CheckFs(elsa.Perms)
			file := args[1]
			val := fs.Mkdir(ctx, file)
			return val
		case Wait:
			fmt.Println(1)
			id := args[1]
			wg.Wait()
			elsa.Recv(id, ctx.Null())
			return ctx.Null()
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

// CheckNet utility to check whether net access is avaliable or not
func CheckNet(perms *options.Perms) {
	if !perms.Net {
		util.LogError("Perms Error: ", "Net is blocked.")
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
		elsa.Recv = func(id quickjs.Value, val quickjs.Value) quickjs.Value {
			result := fn.Call(id, val)
			// defer result.Free()
			return result
		}
		return ctx.Null()
	}
}

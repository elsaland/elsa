package dev

import (
	"fmt"
	"runtime"

	"github.com/elsaland/elsa/core/options"
	"github.com/elsaland/elsa/util"

	"github.com/elsaland/elsa/core"
	"github.com/elsaland/quickjs"
)

func Compile(source string, sourceFile string, fn func(val quickjs.Value), flags *options.Perms, args []string) {
	data, err := core.Asset("typescript/typescript.js")
	if err != nil {
		panic("Asset was not found.")
	}

	runtime.LockOSThread()
	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	context := jsruntime.NewContext()
	defer context.Free()

	core.PrepareRuntimeContext(context, jsruntime, args, flags, "dev")

	globals := context.Globals()
	report := func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		fn(args[0])
		return ctx.Null()
	}
	d := func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		asset, er := core.Asset(args[0].String())
		if er != nil {
			panic("Asset was not found.")
		}
		return ctx.String(string(asset))
	}
	globals.Set("Report", context.Function(report))
	globals.Set("Asset", context.Function(d))
	bundle := string(data) + jsCheck(source, sourceFile)
	result, err := context.Eval(bundle)
	util.Check(err)
	defer result.Free()
}

func jsCheck(source, sourceFile string) string {
	return fmt.Sprintf("typeCheck(`%s`, `%s`);", sourceFile, source)
}

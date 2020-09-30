package dev

import (
	"fmt"
	"runtime"

	"github.com/elsaland/elsa/cmd"
	"github.com/elsaland/elsa/core"
	"github.com/elsaland/quickjs"
)

func Compile(source string, fn func(val quickjs.Value), flags cmd.Perms, args []string) {
	data, err := core.Asset("typescript/typescript.js")
	if err != nil {
		panic("Asset was not found.")
	}
	dts, er := core.Asset("typescript/lib.es6.d.ts")
	if er != nil {
		panic("Asset was not found.")
	}

	runtime.LockOSThread()
	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	context := jsruntime.NewContext()
	defer context.Free()

	core.PrepareRuntimeContext(context, jsruntime, flags, args)

	globals := context.Globals()
	report := func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		fn(args[0])
		return ctx.Null()
	}
	d := func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		return ctx.String(string(dts))
	}
	globals.Set("__report", context.Function(report))
	globals.Set("__getDTS", context.Function(d))
	bundle := string(data) + jsCheck(source)
	result, err := context.Eval(bundle)
	core.Check(err)
	defer result.Free()
}

func jsCheck(source string) string {
	return fmt.Sprintf("ee.emitEvent('typecheck', [`%s`]);", source)
}

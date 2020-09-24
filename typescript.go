package main

import (
	"fmt"
	"runtime"

	"github.com/elsaland/elsa/core"
	"github.com/lithdew/quickjs"
)

func Compile(source string, fn func(val quickjs.Value)) {
	data, err := core.Asset("typescript/typescript.js")
	if err != nil {
		panic("Asset was not found.")
	}
	elsaTSIntegration, err := core.Asset("typescript/elsa.js")
	if err != nil {
		panic("Asset was not found.")
	}
	elsaEvt, err := core.Asset("target/elsa.js")
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
	result, err := context.Eval(string(data))
	defer result.Free()
	core.Check(err)
	result, err = context.Eval(string(elsaEvt))
	defer result.Free()
	core.Check(err)
	result, err = context.Eval(string(elsaTSIntegration))
	defer result.Free()
	core.Check(err)
	result, err = context.Eval(jsCheck(source))
	defer result.Free()
	core.Check(err)
}

func jsCheck(source string) string {
	return fmt.Sprintf("ee.emitEvent('typecheck', [`%s`]);", source)
}

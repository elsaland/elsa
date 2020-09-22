package main

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/elsaland/elsa/cmd"
	"github.com/lithdew/quickjs"
)

func main() {
	runtime.LockOSThread()
	cmd.Execute(cmd.Elsa{Run: run, Bundle: BundleModule})
}

func run(source string, flags cmd.Perms) {
	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	context := jsruntime.NewContext()
	defer context.Free()

	globals := context.Globals()

	globals.Set("__dispatch", context.Function(ElsaNS(flags)))

	snap, _ := Asset("target/elsa.js")

	k, e := context.Eval(string(snap))
	Check(e)
	defer k.Free()

	bundle := BundleModule(source)

	result, e := context.EvalFile(bundle, source)

	defer result.Free()
	if e != nil {
		var evalErr *quickjs.Error
		if errors.As(e, &evalErr) {
			fmt.Println(evalErr.Cause)
			fmt.Println(evalErr.Stack)
		}
		panic(e)
	}
}

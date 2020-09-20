package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"runtime"

	"github.com/divy-work/done/cmd"
	"github.com/lithdew/quickjs"
)

func Check(err error) {
	if err != nil {
		var evalErr *quickjs.Error
		if errors.As(err, &evalErr) {
			fmt.Println(evalErr.Cause)
			fmt.Println(evalErr.Stack)
		}
		panic(err)
	}
}

func main() {
	runtime.LockOSThread()
	cmd.Execute(run)
}

func run(source string) {
	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	context := jsruntime.NewContext()
	defer context.Free()

	globals := context.Globals()

	globals.Set("__dispatch", context.Function(DoneNS))

	snap, _ := Asset("target/done.js")

	k, e := context.Eval(string(snap))
	Check(e)
	defer k.Free()

	bundle := BundleModule(source)
	a := func(val quickjs.Value) {
		ReportDiagnostics(val)
	}
	dat, e := ioutil.ReadFile(source)
	if e != nil {
		panic(e)
	}
	Compile(string(dat), a)
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

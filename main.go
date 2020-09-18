package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/lithdew/quickjs"
)

func check(err error) {
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
	source := os.Args[1:][0]

	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	context := jsruntime.NewContext()
	defer context.Free()
	globals := context.Globals()

	globals.Set("__dispatch", context.Function(DoneNS))

	k, e := context.Eval(NSInject())
	check(e)
	defer k.Free()

	bundle := api.Build(api.BuildOptions{
		EntryPoints: []string{source},
		Outfile:     "output.js",
		Bundle:      true,
		Target:      api.ESNext,
		LogLevel:    api.LogLevelInfo,
	})

	result, e := context.EvalFile(string(bundle.OutputFiles[0].Contents[:]), "s")

	defer result.Free()
	if e != nil {
		var evalErr *quickjs.Error
		if errors.As(e, &evalErr) {
			fmt.Println(evalErr.Cause)
			fmt.Println(evalErr.Stack)
		}
		panic(e)
	}
	fmt.Println()
}

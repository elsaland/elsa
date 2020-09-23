package core

import (
	"errors"
	"fmt"

	"github.com/elsaland/elsa/cmd"
	"github.com/lithdew/quickjs"
)

func Run(source string, bundle string, flags cmd.Perms) {
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

package core

import (
	"errors"
	"fmt"
	"io"

	"github.com/elsaland/elsa/cmd"
	"github.com/elsaland/quickjs"
)

func Run(source string, bundle string, flags cmd.Perms) {
	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	context := jsruntime.NewContext()
	defer context.Free()

	globals := context.Globals()

	globals.Set("__dispatch", context.Function(ElsaNS(flags)))

	snap, _ := Asset("target/elsa.js")

	k, err := context.Eval(string(snap))
	Check(err)
	defer k.Free()

	result, e := context.EvalFile(bundle, source)

	for {
		_, err = jsruntime.ExecutePendingJob()
		if err == io.EOF {
			err = nil
			break
		}
		Check(err)
	}

	defer result.Free()
	if err != nil {
		var evalErr *quickjs.Error
		if errors.As(e, &evalErr) {
			fmt.Println(evalErr.Cause)
			fmt.Println(evalErr.Stack)
		}
		Panic(e)
	}
}

package core

import (
	"errors"
	"fmt"
	"io"

	"github.com/elsaland/elsa/cmd"
	"github.com/elsaland/quickjs"
)

type Recv func(id quickjs.Value, val quickjs.Value)
type Elsa struct {
	flags cmd.Perms
	recv  Recv
}

func Run(source string, bundle string, flags cmd.Perms) {
	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	context := jsruntime.NewContext()
	defer context.Free()

	elsa := &Elsa{flags: flags}

	globals := context.Globals()

	globals.SetFunction("__send", ElsaSendNS(elsa))
	globals.SetFunction("__recv", ElsaRecvNS(elsa))

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
		panic(e)
	}
}

package core

import (
  "github.com/elsaland/elsa/cmd"
  "github.com/elsaland/quickjs"
  "io"
)

type Recv func(id quickjs.Value, val quickjs.Value)
type Elsa struct {
	perms cmd.Perms
	recv  Recv
}

func Run(source string, bundle string, flags cmd.Perms) {
	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	context := jsruntime.NewContext()
	defer context.Free()

	elsa := &Elsa{perms: flags}

	globals := context.Globals()

	globals.SetFunction("__send", ElsaSendNS(elsa))
	globals.SetFunction("__recv", ElsaRecvNS(elsa))

	snap, _ := Asset("target/elsa.js")

	k, err := context.Eval(string(snap))
	Check(err)
	defer k.Free()

	result, err := context.EvalFile(bundle, source)
	Check(err)
  defer result.Free()

	for {
		_, err = jsruntime.ExecutePendingJob()
		if err == io.EOF {
			err = nil
			break
		}
		Check(err)
	}

	if result.IsException() {
	  err = context.Exception()
	  Check(err)
  }
}

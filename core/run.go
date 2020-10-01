package core

import (
	"io"

	"github.com/elsaland/elsa/cmd"
	"github.com/elsaland/quickjs"
)

type Recv func(id quickjs.Value, val quickjs.Value)
type Elsa struct {
	Perms cmd.Perms
	Recv  Recv
}

func PrepareRuntimeContext(cxt *quickjs.Context, jsruntime quickjs.Runtime, flags cmd.Perms, args []string) {

	elsa := &Elsa{Perms: flags}

	globals := cxt.Globals()

	globals.SetFunction("__send", ElsaSendNS(elsa))
	globals.SetFunction("__recv", ElsaRecvNS(elsa))

	snap, _ := Asset("target/elsa.js")

	k, err := cxt.Eval(string(snap))
	Check(err)
	defer k.Free()

	ns := globals.Get("Elsa")
	defer ns.Free()

	__args := cxt.Array()
	for i, arg := range args {
		__arg := cxt.String(arg)
		__args.SetByUint32(uint32(i), __arg)
	}
	ns.Set("args", __args)

	result, err := cxt.EvalFile(bundle, source)
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

}

func Run(source string, bundle string, flags cmd.Perms, args []string) {
	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	cxt := jsruntime.NewContext()
	defer cxt.Free()
	PrepareRuntimeContext(cxt, jsruntime, flags, args)
	result, err := cxt.EvalFile(bundle, source)
	Check(err)
	defer result.Free()

	if result.IsException() {
		err = cxt.Exception()
		Check(err)
	}
}

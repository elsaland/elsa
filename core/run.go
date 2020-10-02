package core

import (
	"io"

	"github.com/elsaland/elsa/core/options"
	"github.com/elsaland/elsa/util"

	"github.com/elsaland/quickjs"
)

func PrepareRuntimeContext(cxt *quickjs.Context, jsruntime quickjs.Runtime, args []string, flags *options.Perms) {
	elsa := &options.Elsa{Perms: flags}

	globals := cxt.Globals()

	globals.SetFunction("__send", ElsaSendNS(elsa))
	globals.SetFunction("__recv", ElsaRecvNS(elsa))

	snap, _ := Asset("target/elsa.js")

	k, err := cxt.Eval(string(snap))
	util.Check(err)
	defer k.Free()

	ns := globals.Get("Elsa")
	defer ns.Free()

	__args := cxt.Array()
	for i, arg := range args {
		__arg := cxt.String(arg)
		__args.SetByUint32(uint32(i), __arg)
	}
	ns.Set("args", __args)

	for {
		_, err = jsruntime.ExecutePendingJob()
		if err == io.EOF {
			err = nil
			break
		}
		util.Check(err)
	}
}

// Run create and dispatch a QuickJS runtime binded with Elsa's OPs configurable using options
func Run(opt options.Options) {
	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	cxt := jsruntime.NewContext()
	defer cxt.Free()

	PrepareRuntimeContext(cxt, jsruntime, opt.Env.Args, opt.Perms)

	result, err := cxt.EvalFile(opt.Source, opt.SourceFile)
	util.Check(err)
	defer result.Free()

	if result.IsException() {
		err = cxt.Exception()
		util.Check(err)
	}

	for {
		_, err = jsruntime.ExecutePendingJob()
		if err == io.EOF {
			err = nil
			break
		}
		util.Check(err)
	}
}

package main

import (
	"io/ioutil"

	"github.com/elsaland/elsa/cmd"
	"github.com/elsaland/elsa/core"
	"github.com/lithdew/quickjs"
)

func TypeCheck(source string) {
	a := func(val quickjs.Value) {
		ReportDiagnostics(val)
	}
	dat, e := ioutil.ReadFile(source)
	if e != nil {
		panic(e)
	}
	Compile(string(dat), a)
}

func RunDev(source string, bundle string, flags cmd.Perms) {
	TypeCheck(source)
	core.Run(source, bundle, flags)
}

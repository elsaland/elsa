package dev

import (
	"github.com/elsaland/elsa/cmd"
	"github.com/elsaland/elsa/core"
	"github.com/elsaland/quickjs"
)

func TypeCheck(source string) {
	a := func(val quickjs.Value) {
		ReportDiagnostics(val)
	}
	Compile(source, a, cmd.Perms{
		Fs: true,
	})
}

func RunDev(source string, bundle string, flags cmd.Perms, args []string) {
	TypeCheck(source)
	core.Run(source, bundle, flags, args)
}

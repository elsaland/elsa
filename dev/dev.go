package dev

import (
	"github.com/elsaland/elsa/cmd"
	"github.com/elsaland/elsa/core"
	"github.com/elsaland/elsa/module"
	"github.com/elsaland/quickjs"
)

// AllowAll allow all flags when in development mode
var AllowAll = cmd.Perms{
	// Allow file system access
	Fs: true,
}

// TypeCheck run typechecking and report the diagnostics
func TypeCheck(source string, args []string) {
	// Callback function for reporting diagnostics to the user
	a := func(val quickjs.Value) {
		ReportDiagnostics(val)
	}
	// Trigger the compiler with the report callback and source
	// allow all perms and specify os args
	Compile(source, a, &AllowAll, args)
}

// RunDev invoke typechecking and execute
func RunDev(source string, bundle string, args []string, config *module.Config) {
	// Run typechecking
	TypeCheck(source, args)
	// Execute bundled script into a quickJS runtime
	core.Run(source, bundle, args, config, &AllowAll)
}

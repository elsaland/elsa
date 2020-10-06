package dev

import (
	"github.com/elsaland/elsa/core"
	"github.com/elsaland/elsa/core/options"
	"github.com/elsaland/quickjs"
)

// AllowAll allow all flags when in development mode
var AllowAll = options.Perms{
	// Allow file system access
	Fs: true,
}

// TypeCheck run typechecking and report the diagnostics
func TypeCheck(source string, sourceFile string, args []string) {
	// Callback function for reporting diagnostics to the user
	a := func(val quickjs.Value) {
		ReportDiagnostics(val)
	}
	// Trigger the compiler with the report callback and source
	// allow all perms and specify os args
	Compile(source, sourceFile, a, &AllowAll, args)
}

// RunDev invoke typechecking and execute
func RunDev(og string, opt options.Options) {
	// Run typechecking
	TypeCheck(og, opt.SourceFile, opt.Env.Args)
	// Execute bundled script into a quickJS runtime
	core.Run(opt)
}

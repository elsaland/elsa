package options

import (
	"github.com/elsaland/quickjs"
)

// Recv callback for an async op
type Recv func(id quickjs.Value, val quickjs.Value) quickjs.Value

// Elsa represents general data for the runtime
type Elsa struct {
	// Permissions
	Perms *Perms
	// Async recv function
	Recv Recv
}

// Environment configure the runtime environment
type Environment struct {
	// Enable or disable color logging
	NoColor bool
	// Command-line args to pass into Elsa.args
	Args []string
	// Whether to run tests associated with `Elsa.tests()`
	RunTests bool
}

// Perms permissions available for Elsa
type Perms struct {
	// File system access
	Fs bool
	// Net access
	Net bool
	// Env access
	Env bool
}

// Options options for dispatching a new Elsa + QuickJS runtime
type Options struct {
	// File name of the source (used for debuging purposes)
	SourceFile string
	// Source code
	Source string
	// Permission
	Perms *Perms
	// Configure Environment
	Env Environment
}

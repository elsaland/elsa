package options

import (
	"github.com/elsaland/quickjs"
)

type Recv func(id quickjs.Value, val quickjs.Value)
type Elsa struct {
	Perms *Perms
	Recv  Recv
}

// Environment configure the runtime environment
type Environment struct {
	// Enable or disable color logging
	NoColor bool
	// Command-line args to pass into Elsa.args
	Args []string
}

type Perms struct {
	Fs bool
}

// Options options for dispatching a new Elsa + QuickJS runtime
type Options struct {
	// File name of the source (used for debuging purposes)
	SourceFile string
	// Source code
	Source string
	// Permission
	Perms *Perms
	// Configure Enviornment
	Env Environment
}

package main

import (
	"github.com/elsaland/elsa/dev"
	"runtime"

	"github.com/elsaland/elsa/bundler"
	"github.com/elsaland/elsa/cmd"
	"github.com/elsaland/elsa/core"
)

func main() {
	runtime.LockOSThread()
	cmd.Execute(cmd.Elsa{
		Run:    core.Run,
		Bundle: bundler.BundleModule,
		Dev:    dev.RunDev,
	})
}

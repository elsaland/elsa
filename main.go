package main

import (
	"runtime"

	"github.com/elsaland/elsa/cmd"
	"github.com/elsaland/elsa/core"
)

func main() {
	runtime.LockOSThread()
	cmd.Execute(cmd.Elsa{Run: core.Run, Bundle: BundleModule})
}

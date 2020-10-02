package testing

import (
	"testing"

	"github.com/elsaland/elsa/cmd"
	"github.com/elsaland/elsa/core"
	"github.com/elsaland/elsa/module"
	"github.com/elsaland/elsa/util"
)

type runTestDesc struct {
	name   string
	source string
}

var RunTestDesc = []runTestDesc{
	{
		"Basic",
		"1 + 1",
	},
	{
		"Bundle local js module",
		"Elsa.readFile('fs/sample.txt')",
	},
}

func TestCore(t *testing.T) {
	// Load config
	config, err := module.ConfigLoad()
	util.Check(err)
	for _, tst := range RunTestDesc {
		// Passing Test
		t.Run(tst.name, func(t *testing.T) {
			// Run the test source with filename as test.js, default config and all perms
			core.Run("test.js", tst.source, []string{}, config, &cmd.Perms{Fs: true})
		})
	}
}

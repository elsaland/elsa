package bundler

import (
	"github.com/elsaland/elsa/module"
	"io/ioutil"
	"testing"

	. "github.com/franela/goblin"
)

type bundleTestDesc struct {
	name     string
	path     string
	category string
}

var TestDesc = []bundleTestDesc{
	{
		"Bundle no-import js module",
		"../testing/basics.js",
		"vanilla",
	},
	{
		"Bundle local js module",
		"../testing/local_imports.js",
		"es",
	},
	{
		"Bundle no-import ts modules",
		"../testing/hello.ts",
		"ts",
	},
}

func readOutData(sourceFile string) string {
	dat, e := ioutil.ReadFile(sourceFile + ".out")
	if e != nil {
		panic(e)
	}
	return string(dat)
}

func TestBundle(t *testing.T) {
	g := Goblin(t)
	g.Describe("Bundle no-import js modules", func() {
		for _, tst := range TestDesc {
			// Passing Test
			g.It(tst.name, func() {
				g.Assert(BundleModule(tst.path)).Equal(readOutData(tst.path))
			})
		}

	})
}

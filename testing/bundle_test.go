package testing

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/elsaland/elsa/bundler"
	"github.com/elsaland/elsa/module"
	"github.com/stretchr/testify/assert"
)

type bundleTestDesc struct {
	name     string
	path     string
	category string
}

var config = &module.Config{}

var TestDesc = []bundleTestDesc{
	{
		"Bundle no-import js module",
		"bundle/basics.js",
		"vanilla",
	},
	{
		"Bundle local js module",
		"bundle/local_imports.js",
		"es",
	},
	{
		"Bundle no-import ts module",
		"bundle/hello.ts",
		"ts",
	},
	{
		"Bundle URL import",
		"bundle/url.ts",
		"url",
	},
}

// utility method to read the expected output for a particular test file
func readOutData(sourceFile string) string {
	dat, e := ioutil.ReadFile(sourceFile + ".out")
	if e != nil {
		panic(e)
	}
	return string(dat)
}

func TestBundle(t *testing.T) {
	for _, tst := range TestDesc {
		// Passing Test
		t.Run(tst.name, func(t *testing.T) {
			assert := assert.New(t)
			// Since URL bundle outputs may differ depending upon the if the import is cached or not
			// therefore, we only check whether it didn't exit with a bad status code.
			// TODO: we might want to do wildcard based assertion
			if tst.category == "url" {
				bundle := bundler.BundleModule(tst.path, false, config)

				assert.NotNil(bundle)
			} else {
				// Remove newlines from the out data and bundle and assert
				bundle := strings.ReplaceAll(strings.ReplaceAll(bundler.BundleModule(tst.path, false, config), "\n", ""), "\r", "")
				expected := strings.ReplaceAll(strings.ReplaceAll(readOutData(tst.path), "\n", ""), "\r", "")

				assert.Equal(bundle, expected)
			}
		})
	}
}

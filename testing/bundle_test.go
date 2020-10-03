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
	minify   bool
}

var config = &module.Config{}

var TestDesc = []bundleTestDesc{
	{
		"Bundle no-import js module",
		"bundle/basics.js",
		"vanilla",
		false,
	},
	{
		"Bundle local js module",
		"bundle/local_imports.js",
		"es",
		false,
	},
	{
		"Bundle no-import ts module",
		"bundle/hello.ts",
		"ts",
		false,
	},
	{
		"Bundle URL import",
		"bundle/url.ts",
		"url",
		false,
	},
	{
		"Bundle minified",
		"bundle/basics.js",
		"vanilla",
		true,
	},
}

// utility method to read the expected output for a particular test file
func readOutData(sourceFile string, minified bool) string {
	path := sourceFile
	if minified {
		path += ".mini.out"
	} else {
		path += ".out"
	}
	dat, e := ioutil.ReadFile(path)
	if e != nil {
		panic(e)
	}
	return string(dat)
}

func TestBundle(t *testing.T) {
	r := strings.NewReplacer("\n", "", "\r", "")

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
				bundle := bundler.BundleModule(tst.path, tst.minify, config)
				expected := readOutData(tst.path, tst.minify)

				// Remove newlines from the out data and bundle and assert
				assert.Equal(r.Replace(bundle), r.Replace(expected))
			}
		})
	}
}

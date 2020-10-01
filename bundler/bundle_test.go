package bundler

import (
	"github.com/elsaland/elsa/module"
	"io/ioutil"
	"testing"
)

var dummy = &module.Config{}

func readOutData(sourceFile string) string {
	dat, e := ioutil.ReadFile(sourceFile + ".out")
	if e != nil {
		panic(e)
	}
	return string(dat)
}

func TestBundleVanillaJS(t *testing.T) {
	if BundleModule("testdata/basics.js", dummy) != readOutData("testdata/basics.js") {
		t.Errorf("Bundling vanilla js module failed.")
	}
}

func TestBundleJSLocalImports(t *testing.T) {
	if BundleModule("testdata/local_imports.js", dummy) != readOutData("testdata/local_imports.js") {
		t.Errorf("Bundling local js module imports failed.")
	}
}

func TestBundleTS(t *testing.T) {
	if BundleModule("testdata/hello.ts", dummy) != readOutData("testdata/hello.ts") {
		t.Errorf("Bundling ts module failed.")
	}
}

package main

import (
	"io/ioutil"
	"testing"
)

func readOutData(sourceFile string) string {
	dat, e := ioutil.ReadFile(sourceFile + ".out")
	if e != nil {
		panic(e)
	}
	return string(dat)
}

func TestBundleVanillaJS(t *testing.T) {
	if BundleModule("testdata/basics.js") != readOutData("testdata/basics.js") {
		t.Errorf("Bundling vanilla js module failed.")
	}
}

func TestBundleJSLocalImports(t *testing.T) {
	if BundleModule("testdata/local_imports.js") != readOutData("testdata/local_imports.js") {
		t.Errorf("Bundling local js module imports failed.")
	}
}

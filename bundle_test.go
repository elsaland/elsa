package main

import (
  "github.com/elsaland/elsa/module"
  "io/ioutil"
  "strings"
  "testing"

  "github.com/elsaland/elsa/bundler"
  . "github.com/franela/goblin"
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
    "testing/bundle/basics.js",
    "vanilla",
  },
  {
    "Bundle local js module",
    "testing/bundle/local_imports.js",
    "es",
  },
  {
    "Bundle no-import ts module",
    "testing/bundle/hello.ts",
    "ts",
  },
  {
    "Bundle URL import",
    "testing/bundle/url.ts",
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
  g := Goblin(t)
  g.Describe("Bundle tests", func() {
    for _, tst := range TestDesc {
      // Passing Test
      g.It(tst.name, func() {
        // Since URL bundle outputs may differ depending upon the if the import is cached or not
        // therefore, we only check whether it didn't exit with a bad status code.
        // TODO: we might want to do wildcard based assertion
        if tst.category == "url" {
          bundle := bundler.BundleModule(tst.path, config)
          g.Assert(bundle)
        } else {
          // Remove newlines from the out data and bundle and assert
          bundle := strings.ReplaceAll(bundler.BundleModule(tst.path, config), "\n", "")
          expected := strings.ReplaceAll(readOutData(tst.path), "\n", "")
          g.Assert(bundle).Equal(expected)
        }
      })
    }
  })
}

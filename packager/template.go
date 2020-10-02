package packager

import (
	"fmt"

	"github.com/elsaland/elsa/module"
)

const source string = `
package main

import (
  "os"
  "github.com/elsaland/elsa/core" 
  "github.com/elsaland/elsa/module"
  "github.com/elsaland/elsa/core/options"
)

func main() {
  snap, _ := Asset("%s")
  toml, _ := Asset("%s")
  config, _ := module.ConfigParse(toml)
  env := options.Environment{
	NoColor: config.Options.NoColor,
	Args:    os.Args[1:],
  }
  opt := options.Options{
	SourceFile: "elsa.js",
	Source:     string(snap),
	Perms:      &options.Perms{Fs: true},
	Env:        env,
  }
  core.Run(opt)
}
`

func GeneratePkgSource(path string) string {
	return fmt.Sprintf(source, path, module.DefaultConfigPath)
}

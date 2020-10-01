package packager

import (
	"fmt"
	"github.com/elsaland/elsa/module"
)

const source string = `package main

import (
  "os"
	"github.com/elsaland/elsa/core" 
	"github.com/elsaland/elsa/cmd" 
  "github.com/elsaland/elsa/module"
)

func main() {
  snap, _ := Asset("%s")
	toml, _ := Asset("%s")
  config, _ := module.ConfigParse(toml)
	core.Run("elsa.js", string(snap), os.Args[1:], config, &cmd.Perms{ Fs: true })
}

`

func GeneratePkgSource(path string) string {
	return fmt.Sprintf(source, path, module.DefaultConfigPath)
}

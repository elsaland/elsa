package packager

import (
	"github.com/elsaland/elsa/module"
	"github.com/elsaland/elsa/util"
	"os"
	"path/filepath"

	"github.com/go-bindata/go-bindata"
)

// PkgSource pack bundled js source into an executable
func PkgSource(source string) {
	c := bindata.NewConfig()

	input := parseInput(source)
	if module.ConfigExists() {
		config := parseInput(module.DefaultConfigPath)
		c.Input = []bindata.InputConfig{input, config}
	} else {
		c.Input = []bindata.InputConfig{input}
	}
	c.Output = "target/elsa-package/asset.go"

	err := bindata.Translate(c)
	util.Check(err)

	entry := GeneratePkgSource(source)
	f, _ := os.Create("target/elsa-package/main.go")

	defer f.Close()
	_, err = f.WriteString(entry)
	util.Check(err)

	ExecBuild("target/elsa-package")
}

func parseInput(path string) bindata.InputConfig {
	return bindata.InputConfig{
		Path:      filepath.Clean(path),
		Recursive: false,
	}
}

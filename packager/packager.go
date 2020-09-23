package packager

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-bindata/go-bindata"
)

// PkgSource pack bundled js source into an executable
func PkgSource(source string) {
	c := bindata.NewConfig()
	input := parseInput(source)
	inputs := []bindata.InputConfig{input}
	c.Input = inputs
	c.Output = "target/elsa-package/asset.go"
	err := bindata.Translate(c)
	if err != nil {
		fmt.Fprintf(os.Stderr, "bindata: %v\n", err)
		os.Exit(1)
	}
	entry := GeneratePkgSource(source)
	f, _ := os.Create("target/elsa-package/main.go")

	defer f.Close()
	f.WriteString(entry)
	ExecBuild("target/elsa-package")
}

func parseInput(path string) bindata.InputConfig {
	return bindata.InputConfig{
		Path:      filepath.Clean(path),
		Recursive: false,
	}
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

package main

import (
	"fmt"
	"os"

	"github.com/evanw/esbuild/pkg/api"
	"github.com/robertkrimen/otto"
)

func main() {
	source := os.Args[1:][0]

	vm := otto.New()

	bundle := api.Build(api.BuildOptions{
		EntryPoints: []string{source},
		Outfile:     "output.js",
		Bundle:      true,
		Target:      api.ES2015,
		Write:       true,
		LogLevel:    api.LogLevelInfo,
	})

	result, _ := vm.Run(string(bundle.OutputFiles[0].Contents[:]))
	fmt.Println(result.String())
	fmt.Println()

}

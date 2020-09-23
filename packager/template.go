package packager

import (
	"fmt"
)

const source string = `
package main

import ( 
	"github.com/elsaland/elsa/core" 
	"github.com/elsaland/elsa/cmd" 
)

func main() {
	snap, _ := Asset("%s")
	core.Run("elsa.js", string(snap), cmd.Perms{ Fs: true })
}
`

func GeneratePkgSource(path string) string {
	return fmt.Sprintf(source, path)
}

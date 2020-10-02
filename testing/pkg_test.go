package testing

import (
	"testing"

	"github.com/elsaland/elsa/packager"
)

func TestPkg(t *testing.T) {
	packager.PkgSource("miscellaneous/args.js")
}

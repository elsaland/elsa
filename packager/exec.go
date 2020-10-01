package packager

import (
	"github.com/elsaland/elsa/util"
	"os/exec"
)

func ExecBuild(path string) {
	cmd := exec.Command("go", "build", ".")
	cmd.Dir = path
	err := cmd.Run()
	util.Check(err)
}

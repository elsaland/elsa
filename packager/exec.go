package packager

import "os/exec"

func ExecBuild(path string) {
	cmd := exec.Command("go", "build", ".")
	cmd.Dir = path
	cmd.Run()
}

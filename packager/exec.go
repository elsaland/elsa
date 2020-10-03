package packager

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func ExecBuild(path string) {
	cmd := exec.Command("go", "build", ".")
	cmd.Dir = path
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		os.Exit(1)
	}
}

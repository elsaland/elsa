package core

import (
	"fmt"

	c "github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
)

var out = colorable.NewColorableStdout()

func LogError(str, msg string) {
	fmt.Fprintln(out, c.Bold(c.Red(str)), msg)
}

func LogInfo(str, extra string) {
	fmt.Fprintln(out, c.Bold(c.Green(str)), extra)
}

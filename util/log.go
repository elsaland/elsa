package util

import (
	"fmt"
	"github.com/fatih/color"
)

func LogError(scope, format string, a ...interface{}) {
	fmt.Fprintf(color.Output, "%s: %s\n", color.New(color.FgRed, color.Bold).Sprint(scope), fmt.Sprintf(format, a...))
}

func LogInfo(scope, format string, a ...interface{}) {
	fmt.Fprintf(color.Output, "%s: %s\n", color.New(color.FgGreen, color.Bold).Sprint(scope), fmt.Sprintf(format, a...))
}

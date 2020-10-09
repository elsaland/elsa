package dev

import (
	"fmt"

	"github.com/elsaland/quickjs"
)

// ReportDiagnostics report TypeScript diagnostics
func ReportDiagnostics(diagnostics quickjs.Value) {
	diag := diagnostics.String()
	if diag != "" {
		fmt.Println(diagnostics.String())
	}
}

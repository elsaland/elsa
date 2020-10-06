package dev

import (
	"fmt"

	"github.com/elsaland/quickjs"
)

func ReportDiagnostics(diagnostics quickjs.Value) {
	fmt.Println(diagnostics.String(), "")
}

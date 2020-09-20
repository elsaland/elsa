package main

import (
	"github.com/lithdew/quickjs"
)

func ReportDiagnostics(diagnostics quickjs.Value) {
	LogError(diagnostics.String(), "")
}

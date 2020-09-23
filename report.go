package main

import (
	"github.com/elsaland/elsa/core"
	"github.com/lithdew/quickjs"
)

func ReportDiagnostics(diagnostics quickjs.Value) {
	core.LogError(diagnostics.String(), "")
}

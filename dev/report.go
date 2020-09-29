package dev

import (
	"github.com/elsaland/elsa/core"
	"github.com/elsaland/quickjs"
)

func ReportDiagnostics(diagnostics quickjs.Value) {
	core.LogError(diagnostics.String(), "")
}

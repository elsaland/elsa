package dev

import (
	"github.com/elsaland/elsa/util"
	"github.com/elsaland/quickjs"
)

func ReportDiagnostics(diagnostics quickjs.Value) {
	util.LogError(diagnostics.String(), "")
}

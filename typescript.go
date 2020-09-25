package main

import (
	"runtime"

	"github.com/elsaland/elsa/core"
	"github.com/elsaland/quickjs"
)

func Compile(source string, fn func(val quickjs.Value)) {
	data, err := core.Asset("typescript/typescript.js")
	if err != nil {
		panic("Asset was not found.")
	}

	dts, er := core.Asset("typescript/lib.es6.d.ts")
	if er != nil {
		panic("Asset was not found.")
	}

	runtime.LockOSThread()

	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	context := jsruntime.NewContext()
	defer context.Free()

	globals := context.Globals()
	report := func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		fn(args[0])
		return ctx.Null()
	}
	d := func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		return ctx.String(string(dts))
	}
	globals.Set("__report", context.Function(report))
	globals.Set("__getDTS", context.Function(d))
	result, err := context.Eval(string(data))
	defer result.Free()
	core.Check(err)
	result, err = context.Eval(jsCheck(source))
	defer result.Free()
	core.Check(err)
}

func jsCheck(source string) string {
	return `
  getDiagnosticsForText(` + "`" + source + "`" + `);
  
  function getDiagnosticsForText(text) {
    const dummyFilePath = "/file.ts";
    const textAst = ts.createSourceFile(dummyFilePath, text, ts.ScriptTarget.ES6);
    const dtsAST = ts.createSourceFile("/lib.es6.d.ts", __getDTS(), ts.ScriptTarget.ES6);
    const files = {[dummyFilePath]: textAst, "/lib.es6.d.ts": dtsAST}
      const options = {};
      const host = {
          fileExists: filePath => files[fileName] != null,
          directoryExists: dirPath => dirPath === "/",
          getCurrentDirectory: () => "/",
      getDirectories: () => [],
          getCanonicalFileName: fileName => fileName,
          getNewLine: () => "\n",
          getDefaultLibFileName: () => "/lib.es6.d.ts",
          getSourceFile: filePath => files[filePath],
          readFile: filePath => filePath === dummyFilePath ? text : undefined,
          useCaseSensitiveFileNames: () => true,
          writeFile: () => {}
      };
      const program = ts.createProgram({
          options,
          rootNames: [dummyFilePath],
          host
      });
  
    let diags = "";
    ts.getPreEmitDiagnostics(program)
    .forEach(diagnostic => {
      if (diagnostic.file) {
        let { line, character } = diagnostic.file.getLineAndCharacterOfPosition(diagnostic.start);
        let message = ts.flattenDiagnosticMessageText(diagnostic.messageText, "\n");
        diags += (diagnostic.file.fileName + " " + (line + 1) + ", " + (character + 1) + ": " + message + "\n");
      } else {
        diags += (ts.flattenDiagnosticMessageText(diagnostic.messageText, "\n"));
      }
      });
    __report(diags)
  }
    `
}

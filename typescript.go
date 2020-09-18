package main

import (
	"runtime"

	"github.com/lithdew/quickjs"
)

func Compile(source string) quickjs.Value {
	data, err := Asset("typescript.js")
	if err != nil {
		panic("Asset was not found.")
	}
	runtime.LockOSThread()

	jsruntime := quickjs.NewRuntime()
	defer jsruntime.Free()

	context := jsruntime.NewContext()
	defer context.Free()

	result, err := context.Eval(string(data))
	defer result.Free()
	Check(err)
	result, err = context.Eval(jsCheck(source))
	defer result.Free()
	Check(err)
	return result
}

func jsCheck(source string) string {
	return (`
	const options = {
		module: ts.ModuleKind.CommonJS,
		target: ts.ScriptTarget.ES2015,
	};
	var source = "` + source + `";
	const compilerHost = ts.createCompilerHost(options);
	const originalGetSourceFile = compilerHost.getSourceFile;
	compilerHost.getSourceFile = (fileName) => {
		console.log(fileName);
		if (fileName === "done.ts") {
			source = typescript.createSourceFile(fileName, source, typescript.ScriptTarget.ES2015, true);
			return source;
		}
		else  return originalGetSourceFile.call(compilerHost, fileName);
	};
	let program = ts.createProgram(["done.ts"], options, compilerHost);
  	let emitResult = program.emit();

	`)
}

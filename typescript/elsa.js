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
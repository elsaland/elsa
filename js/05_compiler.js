var ee = new EventEmitter();
ee.defineEvents(['typecheck']);

ee.addListener('typecheck', (x) => {
  console.log(x)
  getDiagnostics(x)
});

function getDiagnostics(text) {
	const dummyFilePath = text;
	const textAst = ts.createSourceFile(dummyFilePath, Elsa.readFile(text), ts.ScriptTarget.ES6);
	const dtsAST = ts.createSourceFile("/lib.es6.d.ts", __getDTS(), ts.ScriptTarget.ES6);
	const files = {[dummyFilePath]: textAst, "/lib.es6.d.ts": dtsAST}
    const options = { allowJs: true };
    const host = {
        fileExists: filePath => files[filePath] || Elsa.readFile(filePath),
        directoryExists: dirPath => files[dirPath] || Elsa.readFile(dirPath),
        getCurrentDirectory: () => "/",
		getDirectories: () => [],
        getCanonicalFileName: fileName => fileName,
        getNewLine: () => "\n",
        getDefaultLibFileName: () => "/lib.es6.d.ts",
        getSourceFile: filePath => files[filePath] || Elsa.readFile(filePath),
        readFile: filePath => files[filePath] || Elsa.readFile(filePath),
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

// Copyright 2020 elsa.land authors. All rights reserved. MIT license.
const IGNORED_DIAGNOSTICS = [
  // TS2306: File 'file:///Users/rld/src/deno/cli/tests/subdir/amd_like.js' is
  // not a module.
  2306,
  // TS1375: 'await' expressions are only allowed at the top level of a file
  // when that file is a module, but this file has no imports or exports.
  // Consider adding an empty 'export {}' to make this file a module.
  1375,
  // TS1103: 'for-await-of' statement is only allowed within an async function
  // or async generator.
  1103,
  // TS2691: An import path cannot end with a '.ts' extension. Consider
  // importing 'bad-module' instead.
  2691,
  // TS5009: Cannot find the common subdirectory path for the input files.
  5009,
  // TS5055: Cannot write file
  // 'http://localhost:4545/cli/tests/subdir/mt_application_x_javascript.j4.js'
  // because it would overwrite input file.
  5055,
  // TypeScript is overly opinionated that only CommonJS modules kinds can
  // support JSON imports.  Allegedly this was fixed in
  // Microsoft/TypeScript#26825 but that doesn't seem to be working here,
  // so we will ignore complaints about this compiler setting.
  5070,
  // TS7016: Could not find a declaration file for module '...'. '...'
  // implicitly has an 'any' type.  This is due to `allowJs` being off by
  // default but importing of a JavaScript module.
  7016,
];

const options = { allowNonTsExtensions: true };

function typeCheck(file, source) {
  const dummyFilePath = file;
  const textAst = ts.createSourceFile(
    dummyFilePath,
    source,
    ts.ScriptTarget.ES6
  );
  const dtsAST = ts.createSourceFile(
    "/lib.es6.d.ts",
    Asset("typescript/lib.es6.d.ts"),
    ts.ScriptTarget.ES6
  );

  const files = { [dummyFilePath]: textAst, "/lib.es6.d.ts": dtsAST };
  const host = {
    fileExists: (filePath) => {
      return files[filePath] != null || Elsa.exists(filePath);
    },
    directoryExists: (dirPath) => dirPath === "/",
    getCurrentDirectory: () => Elsa.cwd(),
    getDirectories: () => [],
    getCanonicalFileName: (fileName) => fileName,
    getNewLine: () => "\n",
    getDefaultLibFileName: () => "/lib.es6.d.ts",
    getSourceFile: (filePath) => {
      if (files[filePath] != null) return files[filePath];
      else {
        return ts.createSourceFile(
          filePath,
          Elsa.readFile(filePath),
          ts.ScriptTarget.ES6
        );
      }
    },
    readFile: (filePath) => {
      return filePath === dummyFilePath ? text : Elsa.readFile(filePath);
    },
    useCaseSensitiveFileNames: () => true,
    writeFile: () => {},
    resolveModuleNames,
  };
  const program = ts.createProgram({
    options,
    rootNames: [dummyFilePath],
    host,
  });

  let diag = ts.getPreEmitDiagnostics(program).filter(function ({ code }) {
    return code != 5023 && !IGNORED_DIAGNOSTICS.includes(code);
  });
  let diags = ts.formatDiagnosticsWithColorAndContext(diag, host);
  Report(diags);
}

function resolveModuleNames(moduleNames, containingFile) {
  const resolvedModules = [];
  for (const moduleName of moduleNames) {
    let fileName = join(containingFile, "..", moduleName);
    if (moduleName.startsWith("https://")) {
      fileName = moduleName.replace("https://", "/tmp/");
    }
    resolvedModules.push({ resolvedFileName: fileName });
  }
  return resolvedModules;
}

// Joins path segments.  Preserves initial "/" and resolves ".." and "."
// Does not support using ".." to go above/outside the root.
// This means that join("foo", "../../bar") will not resolve to "../bar"
function join(/* path segments */) {
  // Split the inputs into a list of path commands.
  var parts = [];
  for (var i = 0, l = arguments.length; i < l; i++) {
    parts = parts.concat(arguments[i].split("/"));
  }
  // Interpret the path commands to get the new resolved path.
  var newParts = [];
  for (i = 0, l = parts.length; i < l; i++) {
    var part = parts[i];
    // Remove leading and trailing slashes
    // Also remove "." segments
    if (!part || part === ".") continue;
    // Interpret ".." to pop the last segment
    if (part === "..") newParts.pop();
    // Push new path segments.
    else newParts.push(part);
  }
  // Preserve the initial slash if there was one.
  if (parts[0] === "") newParts.unshift("");
  // Turn back into a single string path.
  return newParts.join("/") || (newParts.length ? "/" : ".");
}

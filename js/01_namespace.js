globalThis.Elsa = {
  readFile: (arg) => {
    return globalThis.__dispatch(__ops.FSRead, arg);
  },
  writeFile: (file, content) => {
    return globalThis.__dispatch(__ops.FSWrite, file, content);
  },
  exists: (arg) => {
    return globalThis.__dispatch(__ops.FSExists, arg);
  },
  cwd: () => {
    return globalThis.__dispatch(__ops.FSCwd);
  },
  runPlugin: (dylib, arg) => {
    return globalThis.__dispatch(__ops.Plugin, dylib, arg);
  },
};

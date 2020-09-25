globalThis.Elsa = {
  readFile: (arg) => {
    return globalThis.__send(__ops.FSRead, arg);
  },
  writeFile: (file, content) => {
    return globalThis.__dispatch(__ops.FSWrite, file, content);
  },
  exists: (arg) => {
    return globalThis.__dispatch(__ops.FSExists, arg);
  },
  runPlugin: (dylib, arg) => {
    return globalThis.__send(__ops.Plugin, dylib, arg);
  },
};

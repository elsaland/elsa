globalThis.Elsa = {
  readFile: (arg) => {
    return globalThis.__dispatch(__ops.FSRead, arg);
  },
  exists: (arg) => {
    return globalThis.__dispatch(__ops.FSExists, arg);
  },
  runPlugin: (dylib, arg) => {
    return globalThis.__dispatch(__ops.Plugin, dylib, arg);
  },
};

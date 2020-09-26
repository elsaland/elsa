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
  stats: (arg) => {
    return JSON.parse(globalThis.__dispatch(__ops.FSStats, arg));
  },
  remove: (arg) => {
    return globalThis.__dispatch(__ops.FSRemove, arg);
  },
  cwd: () => {
    return globalThis.__dispatch(__ops.FSCwd);
  },
  runPlugin: (dylib, arg) => {
    return globalThis.__dispatch(__ops.Plugin, dylib, arg);
  },
};

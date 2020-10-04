globalThis.Elsa = {
  readFile: (arg) => {
    return globalThis.__send(__ops.FSRead, arg);
  },
  writeFile: (file, content) => {
    return globalThis.__send(__ops.FSWrite, file, content);
  },
  exists: (arg) => {
    return globalThis.__send(__ops.FSExists, arg);
  },
  stats: (arg) => {
    return JSON.parse(globalThis.__send(__ops.FSStats, arg));
  },
  serve: async function (host, cb) {
      globalThis.__sendAsync(__ops.Serve, function (data) {
        cb(JSON.parse(data))
      }, host);
  },
  remove: (arg) => {
    return globalThis.__send(__ops.FSRemove, arg);
  },
  cwd: () => {
    return globalThis.__send(__ops.FSCwd);
  },
  runPlugin: (dylib, arg) => {
    return globalThis.__send(__ops.Plugin, dylib, arg);
  },
  mkdir: (arg) => {
    return globalThis.__send(__ops.FSMkdir, arg);
  },
};

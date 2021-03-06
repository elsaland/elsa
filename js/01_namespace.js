// Copyright 2020 elsa.land authors. All rights reserved. MIT license.
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
  stat: (arg) => {
    return JSON.parse(globalThis.__send(__ops.FSStat, arg));
  },
  serve: async function (host, cb) {
    globalThis.__sendAsync(
      __ops.Serve,
      function (data) {
        return JSON.stringify(cb(JSON.parse(data)));
      },
      host
    );
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
  env: {
    get: (arg) => {
      return globalThis.__send(__ops.Env, arg);
    },
    set: (env, val) => {
      return globalThis.__send(__ops.Env, env, val);
    },
    toObject: () => {
      return JSON.parse(globalThis.__send(__ops.Env, true));
    },
  },
  *walk(path) {
    const files = JSON.parse(globalThis.__send(__ops.Walk, path));

    for (const file of files) {
      yield file;
    }
  },
};

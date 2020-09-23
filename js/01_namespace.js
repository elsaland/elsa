
globalThis.Elsa = {
    readFile: (arg) => {
        return globalThis.__dispatch(__ops.FSRead, arg);
    },
    runPlugin: (dylib, arg) => {
        return globalThis.__dispatch(__ops.Plugin, dylib, arg);
    }
};


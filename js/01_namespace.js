
globalThis.Elsa = {
    readFile: (arg) => {
        return globalThis.__dispatch(__ops.FSRead, arg);
    }
};



globalThis.Done = {
    readFile: (arg) => {
        return globalThis.__dispatch(__ops.FSRead, arg);
    }
};

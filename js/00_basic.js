
globalThis.Done = {
    readFile: (arg) => {
        return globalThis.__dispatch(1, arg);
    }
};

globalThis.console = {
    trace: (...args) => {
        globalThis.__dispatch(3, ...args);
    },
    debug: (...args) => {
        globalThis.__dispatch(3, ...args);
    },
    log: (...args) => {
        globalThis.__dispatch(3, ...args);
    },
    info: (...args) => {
        globalThis.__dispatch(3, ...args);
    },
    warn: (...args) => {
        globalThis.__dispatch(3, ...args);
    },
    error: (...args) => {
        globalThis.__dispatch(3, ...args);
    },
};

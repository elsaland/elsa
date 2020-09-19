
globalThis.Done = {
    readFile: (arg) => {
        return globalThis.__dispatch("readFile", arg);
    }
};

globalThis.console = {
    trace: (...args) => {
        globalThis.__dispatch("console", ...args);
    },
    debug: (...args) => {
        globalThis.__dispatch("console", ...args);
    },
    log: (...args) => {
        globalThis.__dispatch("console", ...args);
    },
    info: (...args) => {
        globalThis.__dispatch("console", ...args);
    },
    warn: (...args) => {
        globalThis.__dispatch("console", ...args);
    },
    error: (...args) => {
        globalThis.__dispatch("console", ...args);
    },
};

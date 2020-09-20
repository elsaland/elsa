// DOM console bindings
globalThis.console = {
    trace: (...args) => {
        globalThis.__dispatch(__ops.Log, ...args);
    },
    debug: (...args) => {
        globalThis.__dispatch(__ops.Log, ...args);
    },
    log: (...args) => {
        globalThis.__dispatch(__ops.Log, ...args);
    },
    info: (...args) => {
        globalThis.__dispatch(__ops.Log, ...args);
    },
    warn: (...args) => {
        globalThis.__dispatch(__ops.Log, ...args);
    },
    error: (...args) => {
        globalThis.__dispatch(__ops.Log, ...args);
    },
};

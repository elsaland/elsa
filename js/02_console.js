// DOM console bindings
globalThis.console = {
  trace: (...args) => {
    globalThis.__send(__ops.Log, JSON.stringify(...args));
  },
  debug: (...args) => {
    globalThis.__send(__ops.Log, JSON.stringify(...args));
  },
  log: (...args) => {
    globalThis.__send(__ops.Log, JSON.stringify(...args));
  },
  info: (...args) => {
    globalThis.__send(__ops.Log, JSON.stringify(...args));
  },
  warn: (...args) => {
    globalThis.__send(__ops.Log, JSON.stringify(...args));
  },
  error: (...args) => {
    globalThis.__send(__ops.Log, JSON.stringify(...args));
  },
};

// DOM console bindings
globalThis.console = {
  trace: (...args) => {
    val = args[0];
    globalThis.__send(__ops.Log, typeof val, typeof val == "object" ? JSON.stringify(val) : val);
  },
  debug: (...args) => {
    val = args[0];
    globalThis.__send(__ops.Log, typeof val, typeof val == "object" ? JSON.stringify(val) : val);
  },
  log: (...args) => {
    val = args[0];
    globalThis.__send(__ops.Log, typeof val, typeof val == "object" ? JSON.stringify(val) : val);
  },
  info: (...args) => {
    val = args[0];
    globalThis.__send(__ops.Log, typeof val, typeof val == "object" ? JSON.stringify(val) : val);
  },
  warn: (...args) => {
    val = args[0];
    globalThis.__send(__ops.Log, typeof val, typeof val == "object" ? JSON.stringify(val) : val);
  },
  error: (...args) => {
    val = args[0];
    globalThis.__send(__ops.Log, typeof val, typeof val == "object" ? JSON.stringify(val) : val);
  },
};

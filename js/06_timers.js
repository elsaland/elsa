globalThis.setTimeout = (fn, ms) => {
  globalThis.__send(__ops.Timers, ms);
};
globalThis.setTimeout = (fn, ms) => {
  globalThis.__sendAsync(__ops.Timers, fn, ms)
};
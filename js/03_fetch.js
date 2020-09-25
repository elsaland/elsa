globalThis.fetch = async function (...args) {
  return globalThis.__dispatchAsync(__ops.Fetch, ...args);
};

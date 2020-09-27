globalThis.fetch = async function (...args) {
  return globalThis.__sendAsync(__ops.Fetch, ...args);
};

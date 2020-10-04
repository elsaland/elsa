globalThis.fetch = async function (url) {
  return globalThis.__sendAsync(__ops.Fetch, false, url);
};

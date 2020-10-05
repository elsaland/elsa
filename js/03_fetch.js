// Copyright 2020 elsa.land authors. All rights reserved. MIT license.
globalThis.fetch = async function (url) {
  return globalThis.__sendAsync(__ops.Fetch, false, url);
};

// Define op codes
const __ops = {
  FSWrite: 1,
  FSRead: 2,

  Log: 3,

  Plugin: 4,

  Fetch: 5,
};

let initialized = false;

const promiseTable = {};
let promiseNextId = 1;

function init() {
  if (initialized) return;
  initialized = true;
  globalThis.__recv(__recvAsync);
}

function __recvAsync(a, b) {
  console.log(b);
  if (!id) return;
  promiseTable[id].resolve(val);
}

async function __sendAsync(op, ...args) {
  init();
  const id = promiseNextId++;
  globalThis.__send(op, ...[id, ...args]);

  let resolve, reject;
  const promise = new Promise((resolve_, reject_) => {
    resolve = resolve_;
    reject = reject_;
  });
  promise.resolve = resolve;
  promise.reject = reject;

  promiseTable[id] = promise;
  const res = await promise;
  if (res.ok) return res;
  else throw new Error("Error");
}

// Define op codes
const __ops = {
  FSWrite: 1,
  FSRead: 2,
  FSExists: 3,
  FSDirExists: 4,
  FSCwd: 5,
  FSStats: 6,
  Serve: 25,
  FSRemove: 7,
  Fetch: 20,
  Log: 10,
  Plugin: 15,
};

((window) => {
  let initialized = false;

  let promiseTable = {};
  let promiseNextId = 1;

  function init() {
    if (initialized) return;
    initialized = true;
    globalThis.__recv(__recvAsync);
  }

  function __recvAsync(id, val) {
    console.log({id, val})
    if (!id) return;
    promiseTable[id].resolve(val)
  }

  async function __sendAsync(op, ...args) {
    init();
    const id = promiseNextId++;

    let resolve, reject;
    const promise = new Promise((resolve_, reject_) => {
      resolve = resolve_;
      reject = reject_;
    });
    promise.resolve = resolve;
    promise.reject = reject;

    promiseTable[id] = promise;

    globalThis.__send(op, ...[id, ...args]);

    const res = await promise;
    if (res.ok) return res;
    else throw new Error("error");
  }

  Object.assign(window, {
    __sendAsync
  })
})(globalThis)


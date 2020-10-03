# `async`

`async` is a module to provide help with aysncronous tasks.

The async module has been _completely_ ported from Deno's standard modules.

- Deferred - Creates a Promise with the `reject` and `resolve` functions.
- Delay - Resolve a Promise after a given amount of milliseconds
- MuxAsyncIterator - The MuxAsyncIterator class multiplexes multiple async iterators into a single
  stream. The class makes an assumption that the final result (the value returned and not
  yielded from the iterator) does not matter. If there is any result, it is
  discarded.
- pooledMap - Transform values from an (async) iterable into another async iterable. The
  transforms are done concurrently, with a max concurrency defined by the
  poolLimit.

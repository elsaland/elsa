# Elsa

<img align="right" src=assets/logo.svg height="120px">

Elsa is a _minimal_, _fast_ and _secure_ runtime for Javascript and Typescript written in Go.

### Features

- URL based imports
- No fs, net access unless specified
- Compliant to web standards
- Supports TypeScript.

### Coming up

- Module caching
- HTTP server, more Web APIs
- Easy installation scripts
- Bundle your script into a single file
- Create a standalone executable for your script

### Benchmarks

Benchmarks for testdata/console.js for Elsa (unreleased), Deno 1.4.1 and Node 14.4.0
```sh
Benchmark #1: deno run ./testdata/console.js
  Time (mean ± σ):      30.3 ms ±   3.9 ms    [User: 18.8 ms, System: 8.9 ms]
  Range (min … max):    25.7 ms …  45.6 ms    88 runs
 
Benchmark #2: ./elsa ./testdata/console.js
  Time (mean ± σ):      13.4 ms ±   4.2 ms    [User: 5.5 ms, System: 5.6 ms]
  Range (min … max):     8.1 ms …  28.0 ms    212 runs
 
Benchmark #3: node testdata/console.js
  Time (mean ± σ):      79.5 ms ±  16.1 ms    [User: 53.8 ms, System: 13.6 ms]
  Range (min … max):    63.5 ms … 135.0 ms    40 runs
 
Summary
  './elsa ./testdata/console.js' ran
    2.27 ± 0.77 times faster than 'deno run ./testdata/console.js'
    5.94 ± 2.23 times faster than 'node testdata/console.js'
```

### Install

Not yet released, [build from source](#build-from-source) instead.

### Build from source

### Getting Started

Try running a simple program:

```typescript
// hello.ts
import { hello } from "https://x.nest.land/arweave-hello@0.0.2/mod.ts";

hello("Elsa")
```

```shell script
elsa hello.ts
```

### Contributing

Start by creating an issue about your feature or bug! Then, create a PR and we'll land it :smile:

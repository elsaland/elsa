# Elsa

[![Travis Status](https://travis-ci.com/elsaland/elsa.svg?branch=master)](https://travis-ci.com/elsaland/elsa) [![Build Status - Badge][]][build status] [![Discord invite][]][discord invite link]

<img align="right" src=assets/logo.svg height="120px">

Elsa is a _minimal_, _fast_ and _secure_ runtime for Javascript and Typescript written in Go.

### Features

- URL based imports
- No fs, net access unless specified
- Compliant to web standards
- Supports TypeScript.
- Module caching
- Bundle your script into a single file using `elsa bundle script.js`
- Create a standalone executable for your script using `elsa pkg script.js`

### Coming up

- HTTP server, more Web APIs
- Easy installation scripts
- Standard modules
- Typechecking support with `dev` subcommand

### Benchmarks

Benchmark data for the master branch is available at `benchmarks/`

Also see [Comparison with Deno and Node](./COMPARISON.md)

### Install

Not yet released, [build from source](#build-from-source) instead.

### Build from source

You will need Go installed on your machine before building.

Install go-bindata using `go get github.com/go-bindata/go-bindata/...`

Clone the repo on your `$GOPATH` and run `make build` to trigger the build process.

### Getting Started

Try running a simple program:

```typescript
// hello.ts
import { hello } from "https://x.nest.land/arweave-hello@0.0.2/mod.ts";

hello("Elsa");
```

```shell script
> elsa run hello.ts
Hello, Elsa
```

### Contributing

Start by creating an issue about your feature or bug! Then, create a PR and we'll land it :smile:

### License

Elsa.land is licensed under MIT License.

[build status - badge]: https://github.com/elsaland/elsa/workflows/Build/badge.svg
[build status]: https://github.com/elsaland/elsa/actions
[discord invite]: https://img.shields.io/discord/757562931725467709?color=697EC4&label=Discord&logo=discord&logoColor=FDFEFE&style=flat-square
[discord invite link]: https://discord.gg/Dw534ZY

# Elsa

[![Travis Status](https://travis-ci.com/elsaland/elsa.svg?branch=master)](https://travis-ci.com/elsaland/elsa) [![Build Status - Badge][]][Build status] [![Discord invite][]][Discord invite link]

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

### Benchmarks

Benchmark data for the master branch is available at `benchmarks/`

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


[Build Status - Badge]: https://github.com/elsaland/elsa/workflows/Build/badge.svg
[Build status]: https://github.com/elsaland/elsa/actions
[Discord invite]: https://img.shields.io/discord/757562931725467709?color=697EC4&label=Discord&logo=discord&logoColor=FDFEFE&style=flat-square
[Discord invite link]: https://discord.gg/Dw534ZY

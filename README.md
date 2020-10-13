# Elsa

[![Travis Status](https://travis-ci.com/elsaland/elsa.svg?branch=master)](https://travis-ci.com/elsaland/elsa) [![Build Status - Badge][]][build status] [![Discord invite][]][discord invite link]

<img align="right" src=assets/logo.svg height="120px">

Elsa is a _minimal_, _fast_ and _secure_ runtime for JavaScript and TypeScript written in Go, leveraging the power from [QuickJS](https://bellard.org/quickjs/).

### Features

- URL based imports.
- No fs, net access unless specified.
- Compliant to web standards.
- Supports TypeScript.
- Module caching.
- Bundle your script into a single file.
- Create a standalone executable for your bundles.

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

### Credits

- [QuickJS](https://bellard.org/quickjs/) - by [Fabrice Bellard](https://bellard.org/) and Charlie Gordon.
- [Esbuild](https://github.com/evanw/esbuild/) - by [Evan Wallace](https://github.com/evanw)

### FAQs

**Where to use Elsa and where not to...**

Elsa is small and embeddable into Go programs, using Quickjs as the backend comes with certain pros and cons -

PROS: 
 * Suitable for CLI applications and shipping them as an executable
 * Native bindings to system level APIs (webview, etc)
 * Hosting discord bots on your raspberry pie :wink:
 * Embeddable inside your Go programs.
 * General purpose scripting

CONS:

* Slower runtime execution than Node and Deno
* Slower backends
* Double GC 


**Why choose QuickJS over V8?**

QuickJS is a small and embeddable JavaScript engine but it lacks V8's JIT for fast JavaScript execution. Although, it doesn't mean you cannot use Elsa on backends and CPU intensive tasks.

QuickJS has a better startup time than V8 so it would be a strong alternative for CLI apps and short-lived runs.

**Looks like a QuickJS wrapper to me?**

Technically, Node and Deno are also V8 wrappers. All do the same job, init engine - init ops - bundle - run. Most people don't realise that implementing native ops is what makes a _runtime_ and not an _interpreter_.

**What's the status of the project?**

It is in it's _very early stages_ of development i.e. nothing should be considered stable. Feel free to take it for a spin though :)

**What does "minimal" actually refer to?**

The goal is to fullfil the _bare minimum_ requirements needed for development of a project. Elsa, although aims to be extendable via _plugins_. A few examples of features that are not likely to be included in Elsa are tools for formatting, linting and analysis.

In short, develop - package - ship

### License

Elsa.land is licensed under MIT License.

[build status - badge]: https://github.com/elsaland/elsa/workflows/Build/badge.svg
[build status]: https://github.com/elsaland/elsa/actions
[discord invite]: https://img.shields.io/discord/757562931725467709?color=697EC4&label=Discord&logo=discord&logoColor=FDFEFE&style=flat-square
[discord invite link]: https://discord.gg/Dw534ZY

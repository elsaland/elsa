## `core`

The directory contains code for Elsa's runtime agnostic.
It comes with a simple Go API for developers to embed Elsa applications into their Go code.

### Usage

```go
package main

import (
    "github.com/elsaland/elsa/core"
    "github.com/elsaland/elsa/core/options"
)

var source string: = "console.log(1)"

func main() {
    env: = options.Environment {
        NoColor: false, // Set true to disable coloured output
        Args: os.Args[1: ],
    }
    opt: = options.Options {
        SourceFile: "file.js",
        Source: source,
        Perms: & options.Perms {
            Fs: false, // Set true to allow file system access
        },
        Env: env,
    }
    core.Run(opt)
}
```

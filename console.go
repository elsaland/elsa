package main

import (
	"fmt"

	"github.com/lithdew/quickjs"
)

func ConsoleJSInject() string {
	return `globalThis.console = {
		trace: (...args) => {
			globalThis.__console_write(...args);
		},
		debug: (...args) => {
			globalThis.__console_write(...args);
		},
		log: (...args) => {
			globalThis.__console_write(...args);
		},
		info: (...args) => {
			globalThis.__console_write(...args);
		},
		warn: (...args) => {
			globalThis.__console_write(...args);
		},
		error: (...args) => {
			globalThis.__console_write(...args);
		},
	};
`
}

func Console(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
	fmt.Println(args[0])
	return ctx.Null()
}

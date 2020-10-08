package ops

import (
	"github.com/elsaland/elsa/util"
	"github.com/elsaland/quickjs"

	"os"
)

// Env elsa handler
func Env(ctx *quickjs.Context, data []quickjs.Value) quickjs.Value {

	// Elsa.env.get
	if len(data) == 2 {
		key := os.Getenv(data[1].String())
		return ctx.String(key)
	}

	// Elsa.env.set
	if len(data) == 3 {
		err := os.Setenv(data[1].String(), data[2].String())
		util.Check(err)

		return ctx.Null()
	}

	return ctx.Null()
}

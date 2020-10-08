package core

import (
	"time"

	"github.com/elsaland/quickjs"
)

//setTimeout bindings to quickjs engine
func SetTimeout(ctx *quickjs.Context, timeout int64) quickjs.Value {
	time.Sleep(time.Millisecond * time.Duration(timeout))

	// TODO return timeout ID, that can be used to clearTimeout()
	return ctx.Null()
}

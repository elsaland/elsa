package ops

import (
	"sync"
	"time"

	"github.com/elsaland/quickjs"
)

//setTimeout bindings to quickjs engine
func SetTimeout(ctx *quickjs.Context, timeout int64, cb func(), wg *sync.WaitGroup) {
	go func() {
		time.Sleep(time.Millisecond * time.Duration(timeout))
		RunThread(func() {
			CallNonBlock(func() {
				cb()
			})
		})
	}()
}

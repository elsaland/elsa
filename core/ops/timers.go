package ops

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/elsaland/quickjs"
)

//setTimeout bindings to quickjs engine
func SetTimeout(ctx *quickjs.Context, timeout int64, cb func(), wg *sync.WaitGroup) {
	defer runtime.UnlockOSThread()
	time.AfterFunc(time.Millisecond*time.Duration(timeout), func() {
		defer wg.Done()
		cb()
		fmt.Println(2)
	})
}

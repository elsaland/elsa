package ops

import (
	"fmt"
	"net/http"

	"github.com/elsaland/quickjs"
)

func Serve(ctx *quickjs.Context, cb func(val quickjs.Value), id quickjs.Value, host quickjs.Value) {
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
			cb(ctx.String("recv back"))
		})
		http.ListenAndServe(host.String(), nil)
	}()
}

package ops

import (
	"fmt"
	"net/http"

	"github.com/elsaland/quickjs"
)

func Serve(ctx *quickjs.Context, cb func(id quickjs.Value, val quickjs.Value), id quickjs.Value, host quickjs.Value) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:]))
		cb(id, ctx.String("recv back"))
	})
	http.ListenAndServe(host.String(), nil)
}

package ops

import (
  "github.com/elsaland/elsa/util"
  "log"

  "github.com/elsaland/quickjs"
  "github.com/imroc/req"
)

func Fetch(ctx *quickjs.Context, url quickjs.Value) quickjs.Value {
  r, err := req.Get(url.String())
  util.Check(err)
  log.Printf("%+v", r) // print info (try it, you may surprise)
  resp, err := r.ToString()
  util.Check(err)
  return ctx.String(resp)
}

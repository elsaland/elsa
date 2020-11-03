package ops

import (
	"encoding/json"
	"log"
	"net"
	"net/url"

	"github.com/elsaland/quickjs"
	"github.com/valyala/fasthttp"
)

// Serve listens for HTTP requests to host and calls callback sequentially on
// every request.
func Serve(ctx *quickjs.Context, callback func(request quickjs.Value) (response string),
	id quickjs.Value, host quickjs.Value) {
	var (
		reqs  = make(chan *fasthttp.RequestCtx)
		resps = make(chan Response)
		errch = make(chan error, 1)
	)
	go func() {
		httpHandler := func(ctx *fasthttp.RequestCtx) {
			reqs <- ctx
			resp := <-resps
			ctx.SetStatusCode(int(resp.Status))
			_, err := ctx.Write([]byte(resp.Body))
			if err != nil {
				log.Fatalf("%v", err)
			}
		}
		errch <- fasthttp.ListenAndServe(host.String(), httpHandler)
		close(reqs)
	}()
	for {
		select {
		case <-errch:
			// TODO: throw the error as an exception to the JS script
			// see https://github.com/elsaland/elsa/issues/75
			break
		case req := <-reqs:
			reqjson, _ := json.Marshal(Request{
				Method:     string(req.Method()),
				URL:        req.URI(),
				Header:     req.Request.Header,
				Path:       string(req.Path()),
				Host:       string(req.Host()),
				QueryArgs:  req.QueryArgs(),
				PostArgs:   req.PostArgs(),
				PostForm:   req.PostArgs(),
				RemoteAddr: req.RemoteAddr(),
				RequestURI: string(req.RequestURI()),
				Referer:    string(req.Referer()),
				UserAgent:  string(req.UserAgent()),
				LocalAddr:  req.LocalAddr(),
				RemoteIP:   req.RemoteIP(),
				LocalIP:    req.LocalIP(),
			})
			respjson := callback(ctx.String(string(reqjson)))
			var resp Response
			err := json.Unmarshal([]byte(respjson), &resp)
			if err != nil {
				log.Fatal(err)
			}
			resps <- resp
		}
	}
}

// Response response returned by callback from js
type Response struct {
	// Status code
	Status int32
	// Body of the response
	Body string
}

type Request struct {
	Method           string
	URL              *fasthttp.URI
	Path             string
	Header           fasthttp.RequestHeader
	ContentLength    int64
	TransferEncoding []string
	Host             string
	Form             url.Values
	PostForm         *fasthttp.Args
	RemoteAddr       net.Addr
	RequestURI       string
	Referer          string
	UserAgent        string
	LocalAddr        net.Addr
	RemoteIP         net.IP
	LocalIP          net.IP
	PostArgs         *fasthttp.Args
	QueryArgs        *fasthttp.Args
}

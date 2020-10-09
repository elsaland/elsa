package ops

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
			ctx.Write(bytes(resp.Body))
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
				Method:           string(req.Method()),
				URL:              req.URI(),
				Header:           req.Request.Header,
				//ContentLength:    req.ContentLength,
				//TransferEncoding: req.TransferEncoding,
				Host:             string(req.Host()),
				//Form:             req.QueryArgs() + req.PostArgs(),
				PostForm:         req.PostArgs(),
				RemoteAddr:       req.RemoteAddr(),
				RequestURI:       req.RequestURI(),
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
	Method string
	URL *fasthttp.URI
	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0


	Header fasthttp.RequestHeader
	ContentLength int64
	TransferEncoding []string


	// For server requests, Host specifies the host on which the
	// URL is sought. For HTTP/1 (per RFC 7230, section 5.4), this
	// is either the value of the "Host" header or the host name
	// given in the URL itself. For HTTP/2, it is the value of the
	// ":authority" pseudo-header field.
	// It may be of the form "host:port". For international domain
	// names, Host may be in Punycode or Unicode form. Use
	// golang.org/x/net/idna to convert it to either format if
	// needed.
	// To prevent DNS rebinding attacks, server Handlers should
	// validate that the Host header has a value for which the
	// Handler considers itself authoritative. The included
	// ServeMux supports patterns registered to particular host
	// names and thus protects its registered Handlers.
	//
	// For client requests, Host optionally overrides the Host
	// header to send. If empty, the Request.Write method uses
	// the value of URL.Host. Host may contain an international
	// domain name.
	Host string

	// Form contains the parsed form data, including both the URL
	// field's query parameters and the PATCH, POST, or PUT form data.
	// This field is only available after ParseForm is called.
	// The HTTP client ignores Form and uses Body instead.
	Form url.Values

	// PostForm contains the parsed form data from PATCH, POST
	// or PUT body parameters.
	//
	// This field is only available after ParseForm is called.
	// The HTTP client ignores PostForm and uses Body instead.
	PostForm url.Values

	// RemoteAddr allows HTTP servers and other software to record
	// the network address that sent the request, usually for
	// logging. This field is not filled in by ReadRequest and
	// has no defined format. The HTTP server in this package
	// sets RemoteAddr to an "IP:port" address before invoking a
	// handler.
	// This field is ignored by the HTTP client.
	RemoteAddr string

	// RequestURI is the unmodified request-target of the
	// Request-Line (RFC 7230, Section 3.1.1) as sent by the client
	// to a server. Usually the URL field should be used instead.
	// It is an error to set this field in an HTTP client request.
	RequestURI string
}

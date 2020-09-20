package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/evanw/esbuild/pkg/api"
)

func BundleModule(source string) string {
	bundle := api.Build(api.BuildOptions{
		EntryPoints: []string{source},
		Outfile:     "output.js",
		Bundle:      true,
		Target:      api.ESNext,
		LogLevel:    api.LogLevelInfo,
		Plugins: []func(api.Plugin){

			func(plugin api.Plugin) {
				plugin.SetName("url-loader")
				plugin.AddResolver(api.ResolverOptions{Filter: "^https?://"},
					func(args api.ResolverArgs) (api.ResolverResult, error) {
						// Get the data
						resp, _ := http.Get(args.Path)
						fileName := buildFileName(args.Path)
						defer resp.Body.Close()
						file, err := ioutil.TempFile("", fileName)
						if err != nil {
							log.Fatal(err)
						}
						io.Copy(file, resp.Body)

						defer file.Close()
						return api.ResolverResult{Path: file.Name(), Namespace: "url-loader"}, nil
					})
				plugin.AddLoader(api.LoaderOptions{Filter: "^", Namespace: "url-loader"},
					func(args api.LoaderArgs) (api.LoaderResult, error) {
						dat, _ := ioutil.ReadFile(args.Path)
						contents := string(dat)
						return api.LoaderResult{Contents: &contents, Loader: api.LoaderTS}, nil
					})

			},
		},
	})
	return string(bundle.OutputFiles[0].Contents[:])
}

func buildFileName(fileURL string) string {
	fileUrl, _ := url.Parse(fileURL)

	path := fileUrl.Path
	segments := strings.Split(path, "/")

	return segments[len(segments)-1]
}

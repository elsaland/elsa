package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/asaskevich/govalidator"
	"github.com/elsaland/elsa/core"
	"github.com/evanw/esbuild/pkg/api"
)

var cache = ElsaCache{os.TempDir()}

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
						possibleCachePath := cache.UrlToPath(args.Path)
						if cache.InCache(possibleCachePath) && cache.Exists(possibleCachePath) {
							return api.ResolverResult{Path: possibleCachePath, Namespace: ""}, nil
						}
						// Get the data
						f := BundleURL(args.Path)
						return api.ResolverResult{Path: f, Namespace: ""}, nil

					})
			},
		},
	})
	return string(bundle.OutputFiles[0].Contents[:])
}

func BundleURL(uri string) string {
	core.LogInfo("Downloading", uri)
	resp, _ := http.Get(uri)
	fileName := cache.BuildFileName(uri)
	defer resp.Body.Close()
	file, err := cache.Create(fileName)
	if err != nil {
		core.LogError("Internal", fmt.Sprintf("%s", err))
		os.Exit(1)
	}
	io.Copy(file, resp.Body)
	defer file.Close()

	api.Build(api.BuildOptions{
		EntryPoints: []string{file.Name()},
		Outfile:     "output.js",
		Bundle:      true,
		Target:      api.ES2015,
		LogLevel:    api.LogLevelInfo,
		Plugins: []func(api.Plugin){

			func(plugin api.Plugin) {
				plugin.SetName("url-loader2")
				plugin.AddResolver(api.ResolverOptions{Filter: ".*?"},
					func(args api.ResolverArgs) (api.ResolverResult, error) {
						dir := filepath.Dir(file.Name())
						possibleCachePath := path.Join(dir, args.Path)
						if cache.InCache(possibleCachePath) && cache.Exists(possibleCachePath) {
							return api.ResolverResult{Path: possibleCachePath, Namespace: ""}, nil
						}
						if govalidator.IsURL(args.Path) {
							bundle := BundleURL(args.Path)
							return api.ResolverResult{Path: bundle, Namespace: ""}, nil
						}
						base, err := url.Parse(uri)

						pth, err := url.Parse(args.Path)
						loc := base.ResolveReference(pth).String()

						core.LogInfo("Downloading", loc)
						// Get the data
						resp, _ := http.Get(loc)
						fileName := cache.BuildFileName(loc)
						defer resp.Body.Close()
						file, err := cache.Create(fileName)
						if err != nil {
							core.LogError("Internal", fmt.Sprintf("%s", err))
							os.Exit(1)
						}
						io.Copy(file, resp.Body)

						defer file.Close()
						core.LogInfo("Downloaded", file.Name())
						return api.ResolverResult{Path: file.Name(), Namespace: ""}, nil

					})
			},
		},
	})
	return file.Name()
}

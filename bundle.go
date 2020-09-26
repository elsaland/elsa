package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
							return api.ResolverResult{Path: possibleCachePath, Namespace: "url-loader"}, nil
						}
						// Get the data
						f := BundleURL(args.Path)
						return api.ResolverResult{Path: f, Namespace: "url-loader"}, nil

					})
				plugin.AddLoader(api.LoaderOptions{Filter: ".*?", Namespace: "url-loader"},
					func(args api.LoaderArgs) (api.LoaderResult, error) {
						p := args.Path
						if cache.InCache(args.Path) && !cache.Exists(args.Path) {
							c := cache.PathToUrl(args.Path)
							resp, _ := http.Get(c)
							fileName := cache.BuildFileName(c)
							defer resp.Body.Close()
							file, err := cache.Create(fileName)
							if err != nil {
								core.LogError("Internal", fmt.Sprintf("%s", err))
								os.Exit(1)
							}
							io.Copy(file, resp.Body)

							defer file.Close()
							p = file.Name()
							core.LogInfo("Downloaded", file.Name())
						}
						core.LogInfo("Loading", p)
						dat, e := ioutil.ReadFile(p)
						if e != nil {
							panic(e)
						}
						contents := string(dat)
						return api.LoaderResult{Contents: &contents, Loader: api.LoaderTS}, nil
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
		Target:      api.ESNext,
		LogLevel:    api.LogLevelInfo,
		Plugins: []func(api.Plugin){

			func(plugin api.Plugin) {
				plugin.SetName("url-loader2")
				plugin.AddResolver(api.ResolverOptions{Filter: ".*?"},
					func(args api.ResolverArgs) (api.ResolverResult, error) {
						dir := filepath.Dir(file.Name())
						possibleCachePath := path.Join(dir, args.Path)
						if cache.InCache(possibleCachePath) && cache.Exists(possibleCachePath) {
							return api.ResolverResult{Path: possibleCachePath, Namespace: "url-loader2"}, nil
						}
						if govalidator.IsURL(args.Path) {
							bundle := BundleURL(args.Path)
							return api.ResolverResult{Path: bundle, Namespace: "url-loader2"}, nil
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
						return api.ResolverResult{Path: file.Name(), Namespace: "url-loader2"}, nil

					})
				plugin.AddLoader(api.LoaderOptions{Filter: ".*?", Namespace: "url-loader2"},
					func(args api.LoaderArgs) (api.LoaderResult, error) {
						p := args.Path
						if cache.InCache(args.Path) && !cache.Exists(args.Path) {
							c := cache.PathToUrl(args.Path)
							resp, _ := http.Get(c)
							fileName := cache.BuildFileName(c)
							defer resp.Body.Close()
							file, err := cache.Create(fileName)
							if err != nil {
								core.LogError("Internal", fmt.Sprintf("%s", err))
								os.Exit(1)
							}
							io.Copy(file, resp.Body)

							defer file.Close()
							p = file.Name()
							core.LogInfo("Downloaded", file.Name())
						}
						core.LogInfo("Loading", p)
						dat, e := ioutil.ReadFile(p)
						if e != nil {
							panic(e)
						}
						contents := string(dat)
						return api.LoaderResult{Contents: &contents, Loader: api.LoaderTS}, nil
					})

			},
		},
	})
	return file.Name()
}

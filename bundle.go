package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/evanw/esbuild/pkg/api"
)

var cache = ElsaCache{ os.TempDir() }

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
						} else {
							LogInfo("Downloading", args.Path)
							// Get the data
							resp, _ := http.Get(args.Path)
							fileName := cache.BuildFileName(args.Path)
							defer resp.Body.Close()
							file, err := cache.Create(fileName)
							if err != nil {
								LogError("Internal", fmt.Sprintf("%s", err))
								os.Exit(1)
							}
							io.Copy(file, resp.Body)

							defer file.Close()
							LogInfo("Downloaded", file.Name())
							return api.ResolverResult{Path: file.Name(), Namespace: "url-loader"}, nil
						}

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
								LogError("Internal", fmt.Sprintf("%s", err))
								os.Exit(1)
							}
							io.Copy(file, resp.Body)

							defer file.Close()
							p = file.Name()
							LogInfo("Downloaded", file.Name())
						}
						LogInfo("Loading", p)
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

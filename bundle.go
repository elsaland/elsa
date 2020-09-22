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
						possibleCachePath := urlToPath(args.Path)
						if inCache(possibleCachePath) && exists(possibleCachePath) {
							return api.ResolverResult{Path: possibleCachePath, Namespace: "url-loader"}, nil
						} else {
							LogInfo("Downloading", args.Path)
							// Get the data
							resp, _ := http.Get(args.Path)
							fileName := buildFileName(args.Path)
							defer resp.Body.Close()
							file, err := create(fileName)
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
						if inCache(args.Path) && !exists(args.Path) {
							c := pathToUrl(args.Path)
							resp, _ := http.Get(c)
							fileName := buildFileName(c)
							defer resp.Body.Close()
							file, err := create(fileName)
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

func buildFileName(fileURL string) string {
	fileUrl, _ := url.Parse(fileURL)
	path := path.Join(os.TempDir(), fileUrl.Host, fileUrl.Path)
	return path
}

func pathToUrl(path string) string {
	parts := strings.Split(path, "/")[2:]
	url, _ := url.Parse("https://" + strings.Join(parts, "/"))
	return url.String()
}

func urlToPath(url string) string {
	parts := strings.Split(url, "//")[1]
	cache := path.Join(os.TempDir(), parts)
	return cache
}

func inCache(path string) bool {
	if strings.HasPrefix(path, os.TempDir()) {
		return true
	}
	return false
}

func exists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

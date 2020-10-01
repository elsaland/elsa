package bundler

import (
  "fmt"
  "github.com/elsaland/elsa/module"
  "github.com/elsaland/elsa/util"
  "io"
  "net/http"
  "net/url"
  "os"
  "path"
  "path/filepath"

  "github.com/asaskevich/govalidator"
  "github.com/evanw/esbuild/pkg/api"
)

var cache = ElsaCache{os.TempDir()}

func BundleModule(path string, config *module.Config) string {
  bundle := api.Build(api.BuildOptions{
    EntryPoints: []string{path},
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
  resp, _ := http.Get(uri)
  fileName := cache.BuildFileName(uri)
  util.LogInfo("Downloading", fmt.Sprintf("%s => %s", uri, fileName))
  defer resp.Body.Close()
  file, err := cache.Create(fileName)
  if err != nil {
    util.LogError("Internal", fmt.Sprintf("%s", err))
    os.Exit(1)
  }
  _, err = io.Copy(file, resp.Body)
  if err != nil {
    util.LogError("Internal", fmt.Sprintf("%s", err))
    os.Exit(1)
  }
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
              uri = args.Path
              cha := make(chan string)
              go func(url string, u chan string) {
                u <- BundleURL(url)
              }(args.Path, cha)
              bundle := <-cha
              return api.ResolverResult{Path: bundle, Namespace: ""}, nil
            }
            base, err := url.Parse(uri)
            pth, err := url.Parse(args.Path)
            loc := base.ResolveReference(pth).String()
            fileName := cache.BuildFileName(loc)
            util.LogInfo("Downloading", fmt.Sprintf("%s => %s", loc, file.Name()))
            // Get the data
            resp, _ := http.Get(loc)
            defer resp.Body.Close()
            file, err := cache.Create(fileName)
            if err != nil {
              util.LogError("Internal", fmt.Sprintf("%s", err))
              os.Exit(1)
            }
            _, err = io.Copy(file, resp.Body)
            if err != nil {
              util.LogError("Internal", fmt.Sprintf("%s", err))
              os.Exit(1)
            }

            defer file.Close()

            return api.ResolverResult{Path: file.Name(), Namespace: ""}, nil

          })
      },
    },
  })
  return file.Name()
}

package main

import "github.com/evanw/esbuild/pkg/api"

func BundleModule(source string) string {
	bundle := api.Build(api.BuildOptions{
		EntryPoints: []string{source},
		Outfile:     "output.js",
		Bundle:      true,
		Target:      api.ESNext,
		LogLevel:    api.LogLevelInfo,
	})
	return string(bundle.OutputFiles[0].Contents[:])
}

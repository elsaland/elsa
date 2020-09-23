// +build linux freebsd macos

package core

import "plugin"

type PluginFunction func(val interface{}) interface{}

func OpenPlugin(path string) plugin.Symbol {
	p, err := plugin.Open(path)
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("ElsaPlugin")
	if err != nil {
		panic(err)
	}
	return f
}

func RunPlugin(symbol plugin.Symbol, arg interface{}) interface{} {
	return symbol.(PluginFunction)(arg)
}

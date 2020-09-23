// +build linux freebsd macos

package core

import "plugin"

func OpenPlugin(path string, arg interface{}) interface{} {
	p, err := plugin.Open(path)
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	*v.(*interface{}) = arg
	f, err := p.Lookup("ElsaPlugin")
	if err != nil {
		panic(err)
	}
	return f.(func() interface{})()
}

//go:build !windows
// +build !windows

package core

import "plugin"

// OpenPlugin open a dynamic lib and call the exported ElsaPlugin function
// with the args provided by the plugin op
// Currently, not compatible with windows
func OpenPlugin(path string, arg interface{}) interface{} {
	// open the plugin
	p, err := plugin.Open(path)
	if err != nil {
		panic(err)
	}
	// lookup for exported variable and assign the argument
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	*v.(*interface{}) = arg
	// lookup for ElsaPlugin export and call the function
	f, err := p.Lookup("ElsaPlugin")
	if err != nil {
		panic(err)
	}
	return f.(func() interface{})()
}

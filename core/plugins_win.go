// +build windows

package core

import "os"

func OpenPlugin(path string, arg interface{}) interface{} {
	LogError("Not supported", "Go plugins are not supported for windows. See https://github.com/golang/go/issues/19282")
	os.Exit(1)
	return nil
}

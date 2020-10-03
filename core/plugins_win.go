// +build windows

package core

import (
	"os"

	"github.com/elsaland/elsa/util"
)

// OpenPlugin Go plugins have not yet been implemented on windows
func OpenPlugin(path string, arg interface{}) interface{} {
	util.LogError("Not supported", "Go plugins are not supported for windows. See https://github.com/golang/go/issues/19282")
	os.Exit(1)
	return nil
}

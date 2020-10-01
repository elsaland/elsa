// +build windows

package core

import (
  "github.com/elsaland/elsa/util"
  "os"
)

func OpenPlugin(path string, arg interface{}) interface{} {
  util.LogError("Not supported", "Go plugins are not supported for windows. See https://github.com/golang/go/issues/19282")
  os.Exit(1)
  return nil
}

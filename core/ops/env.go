package ops

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/elsaland/elsa/util"
	"github.com/elsaland/quickjs"
)

// Env elsa handler
func Env(ctx *quickjs.Context, data []quickjs.Value) quickjs.Value {

	// Elsa.env.get
	if len(data) == 2 && data[1].IsString() {
		key := os.Getenv(data[1].String())
		return ctx.String(key)
	}

	// Elsa.env.set
	if len(data) == 3 && (data[1].IsString() && data[2].IsString()) {
		err := os.Setenv(data[1].String(), data[2].String())
		util.Check(err)

		return ctx.Null()
	}

	// Elsa.env.toObject
	if len(data) == 2 && data[1].IsBool() {
		getenvironment := func(envs []string, getkeyval func(item string) (key, val string)) map[string]string {
			items := make(map[string]string)
			for _, item := range envs {
				key, val := getkeyval(item)
				items[key] = val
			}
			return items
		}

		// get envs map
		environment := getenvironment(os.Environ(), func(item string) (key, val string) {
			splits := strings.Split(item, "=")
			key = splits[0]
			val = splits[1]
			return
		})

		jsonEnv, err := json.Marshal(environment)

		util.Check(err)
		// convert to json string object
		return ctx.String(string(jsonEnv))
	}

	return ctx.Null()
}

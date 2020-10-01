package core

import (
	"encoding/json"
	"fmt"
	"github.com/elsaland/elsa/util"
	"testing"
)

const testString = `"This is a string"`
const testArray = `["string", 100, {}]`
const testNumber = `120`
const testDiverseJSON = `{
	"str": "foo",
	"num": 100,
	"bool": false,
	"null": null,
	"array": ["foo", "bar", "baz"],
	"obj": { "a": 1, "b": 2 }
  }`

func expectPass(str string, t *testing.T) {
	var result interface{}
	err := json.Unmarshal([]byte(str), &result)
	util.Check(err)
	prty, err := Marshal(result)
	util.Check(err)
	fmt.Println(string(prty))
}

func TestString(t *testing.T)      { expectPass(testString, t) }
func TestDiverseJSON(t *testing.T) { expectPass(testDiverseJSON, t) }
func TestNumber(t *testing.T)      { expectPass(testNumber, t) }
func TestArray(t *testing.T)       { expectPass(testArray, t) }

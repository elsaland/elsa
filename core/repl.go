package core

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/elsaland/quickjs"
)

func Repl() {
	stringToEval := ""
	fmt.Println("Elsa REPL")
	fmt.Println("exit using ctrl+c or close()")

	for true {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		stringToEval += text

		fmt.Println(Eval(stringToEval))
	}
}

func Eval(text string) string {
	// repl close function
	closeEval := func(ctx *quickjs.Context, this quickjs.Value, args []quickjs.Value) quickjs.Value {
		os.Exit(1)
		return ctx.Null()
	}

	evalRuntime := quickjs.NewRuntime()
	defer evalRuntime.Free()

	evalContext := evalRuntime.NewContext()
	defer evalContext.Free()

	//TODO(buttercubz) set globals functions
	globalsEval := evalContext.Globals()

	globalsEval.Set("close", evalContext.Function(closeEval))

	result, err := evalContext.Eval(text)
	check(err)
	defer result.Free()

	return result.String()
}

// check errors without exit
func check(err error) {
	if err != nil {
		var evalErr *quickjs.Error
		if errors.As(err, &evalErr) {
			fmt.Println(evalErr.Cause)
			fmt.Println(evalErr.Stack)
		}
	}
}

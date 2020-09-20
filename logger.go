package main

import (
	"fmt"

	c "github.com/logrusorgru/aurora"
)

func LogError(str string) {
	fmt.Println(c.Bold(c.Red(str)))
}

func LogInfo(str, extra string) {
	fmt.Println(c.Bold(c.Green(str)), extra)
}

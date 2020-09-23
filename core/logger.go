package core

import (
	"fmt"

	c "github.com/logrusorgru/aurora"
)

func LogError(str, msg string) {
	fmt.Println(c.Bold(c.Red(str)), msg)
}

func LogInfo(str, extra string) {
	fmt.Println(c.Bold(c.Green(str)), extra)
}

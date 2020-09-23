package main

import "fmt"

var V interface{}

func ElsaPlugin() interface{} {
	fmt.Printf("Hello, %s\n", V)
	return "Some returned data"
}

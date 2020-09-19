package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"sort"
)

func main() {
	files, _ := ioutil.ReadDir("js/")

	sort.Sort(ByNumericalFilename(files))

	binCmd := []string{"-o", "data.go", "typescript/", "js/"}
	var finalSource string
	for _, f := range files {
		log.Printf("Bundling %s\n", f.Name())
		source, _ := ioutil.ReadFile("js/" + f.Name())
		finalSource += string(source)
	}
	ioutil.WriteFile("target/done.js", []byte(finalSource), 0644)
	cmd := exec.Command("go-bindata", binCmd...)
	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
}

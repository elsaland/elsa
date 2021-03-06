package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/js"
)

func main() {
	files, _ := ioutil.ReadDir("js")

	sort.Sort(ByNumericalFilename(files))

	m := minify.New()
	m.AddFunc("text/javascript", js.Minify)

	binCmd := []string{"run", "github.com/go-bindata/go-bindata/go-bindata", "-pkg", "core", "-o", "./core/data.go", "typescript/", "target/"}
	var finalSource string
	for _, f := range files {
		log.Printf("Bundling %s\n", f.Name())

		file, err := os.Open(filepath.Join("js", f.Name()))
		if err != nil {
			log.Fatalf("Got error opening %s: %v", f.Name(), err)
		}

		buf := new(bytes.Buffer)
		if err := m.Minify("text/javascript", buf, file); err != nil {
			log.Fatalf("Got error minifying %s: %v", f.Name(), err)
		}

		finalSource += buf.String() + "\n"
	}
	err := os.Mkdir("target", 0750)
	if err != nil {
		log.Fatalf("Error in making directory - %v", err)
	}
	err = ioutil.WriteFile(filepath.Join("target", "elsa.js"), []byte(finalSource), 0644)
	if err != nil {
		log.Fatalf("Error writing file %v", err)
	}
	cmd := exec.Command("go", binCmd...)
	log.Printf("Running command and waiting for it to finish...")
	err = cmd.Run()
	log.Printf("Command finished with error: %v", err)
}

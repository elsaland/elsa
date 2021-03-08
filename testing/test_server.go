package testing

import (
	"log"
	"net/http"
)

// StartTestServer start a new local test server for bundler tests.
func StartTestServer() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8100", nil))
}

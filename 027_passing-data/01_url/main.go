package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)

	// Not found handler for code 404
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// retrieve the parameter send from GET request url.
	// req.FormValue
	v := req.FormValue("q")
	fmt.Fprintln(w, "Do my search: "+v)
}

// visit this page:
// http://localhost:8080/?q=dog

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

// initial template
// func ParseGlob(pattern string) (*Template, error)
// template.ParseGlob() use filepath.Glob() function
// read all the file content which glob matched
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	// process form submission here
	// http.Redirect
	// func Redirect(w ResponseWriter, r *Request, url string, code int)
	http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)

	// template.ExecuteTemplate() render the template by pass in the
	// variable interface that is last one parameter.
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

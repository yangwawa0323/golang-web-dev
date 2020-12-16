package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

// initial function of golang code
func init() {

	/*
		func Must(t *Template, err error) *Template {
			if err != nil {
				panic(err)
			}
			return t
		}
		template.Must() function capture the error and return the *Template, it is embeded function
	*/
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	// body
	// make the []bytes with length of http.Request.ContentLength.
	bs := make([]byte, req.ContentLength)
	// from request body read the content to the slice `bs`
	req.Body.Read(bs)
	// convert the slice bytes to string.
	body := string(bs)

	// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	// ExecuteTemplate() function read data from `data interface{}` and render to `index.gohtml`
	// finally write to http.Response
	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

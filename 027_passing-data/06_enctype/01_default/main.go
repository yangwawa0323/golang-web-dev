package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// initial a person
// p := person{FirstName: "Yang", LastName: "Kun", Subscribed: false ,}
// p2 := person{"Yang","Kun",true}
// pr := &person{FirstName:"Yang" , LastName:"Kun" , Subscribed: false , }
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
	// `make` function
	// func make(t Type, size ...IntegerType) Type
	// http.Request.ContentLength
	// build a slice with length
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	io.WriteString(os.Stdout, body)
	// template.ExecuteTemplate(wr io.Writer, name string, data interface{})
	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

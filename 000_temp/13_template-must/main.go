package main

import (
	"html/template"
	"io"
	"os"
	"strings"
)

type message struct {
	Context map[string]interface{}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	msg1 := &message{}
	msg1.Context = make(map[string]interface{})
	msg1.Context["Body"] = "Index page"
	tpl.ExecuteTemplate(os.Stdout, "index.gohtml", msg1)
	io.WriteString(os.Stdout, strings.Repeat("\n================\n", 3))
	msg2 := message{Context: map[string]interface{}{"Body": "About me"}}
	tpl.ExecuteTemplate(os.Stdout, "about.gohtml", &msg2)
}

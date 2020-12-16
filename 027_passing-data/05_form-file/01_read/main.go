package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {

	var s string
	fmt.Println(req.Method)
	if req.Method == http.MethodPost {

		// open
		// FormFile returns the first file for the provided form key.
		// FormFile calls ParseMultipartForm and ParseForm if necessary.
		// http.Request.FormFile()
		// func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error)

		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// for your information
		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		// read
		// ioutil.ReadAll(f)
		//func ReadAll(r io.Reader) ([]byte, error)
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// ioutil.ReadAll() function return []byte, err,
		// You need convert []byte to string format
		s = string(bs)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}

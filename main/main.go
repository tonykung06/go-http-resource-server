package main

import (
	// "fmt"
	"bufio"
	// "io/ioutil"
	"net/http"
	"os"
	"strings"
)

//the simplest way to register a http handler
func useHttpHandlerFunc() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`Hello World`))
	})
}

func useHttpHandlerObject() {
	http.Handle("/", new(MyHandler))
}

func main() {
	// useHttpHandlerFunc()
	useHttpHandlerObject()
	http.ListenAndServe(":8000", nil)
}

type MyHandler struct {
	http.Handler //will be implementing this interface, actually this explicit statement is not required
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public/" + req.URL.Path
	//reading whole files into memory, OMG!
	// data, err := ioutil.ReadFile(string(path))

	f, err := os.Open(path)

	if err == nil {
		bufferedReader := bufio.NewReader(f)

		var contentType string
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".mp4") {
			contentType = "video/mp4"
		} else {
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)
		// w.Write(data)
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}

package main

import (
	// "fmt"
	"bufio"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

//the simplest way to register a http handler
func useHttpHandlerFunc() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`Hello World`))
	})
	http.ListenAndServe(":8000", nil)
}

func useFileServer() {
	http.ListenAndServe(":8000", http.FileServer(http.Dir("public")))
}

func useHttpHandlerObject() {
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":8000", nil)
}

func useBufferedHttpHandlerObject() {
	http.Handle("/", new(MyBufferedIoHandler))
	http.ListenAndServe(":8000", nil)
}

type MyBufferedIoHandler struct {
	http.Handler //will be implementing this interface, actually this explicit statement is not required
}

type MyHandler struct {
	http.Handler //will be implementing this interface, actually this explicit statement is not required
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public/" + req.URL.Path
	//reading whole files into memory, OMG!
	data, err := ioutil.ReadFile(string(path))
	if err == nil {
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
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}

func (this *MyBufferedIoHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "public/" + req.URL.Path
	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
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
		bufferedReader.WriteTo(w)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}

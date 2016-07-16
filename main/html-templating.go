package main

import (
	"net/http"
	"text/template"
)

func useBasicHtmlTemplate() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		tmpls := template.New("tmpl")
		tmpls.New("test").Parse(doc)
		tmpls.New("firstPart").Parse(firstPart)
		context := Context{
			"the message",
			req.URL.Path,
			[3]string{"Apple", "Orange", "Banana"},
		}
		tmpls.Lookup("test").Execute(w, context)
	})

	http.ListenAndServe(":8000", nil)
}

const doc = `
<!DOCTYPE html>
<html>
	<head>
		<title>html template example title</title>
	</head>
	<body>
		{{template "firstPart" .}}
		<h1>List of fruits</h1>
		<ul>
			{{range .Fruits}}
				<li>{{.}}</li>
			{{else}}
				<li>No fruits</li>
			{{end}}
		</ul>
	</body>
</html>
`

const firstPart = `
	<h1>Hello {{.Message}}</h1>
	{{if eq .Path "/hello"}}
		<p>Hello World<p>
	{{else}}
		<p>WHAT???<p>
	{{end}}
`

type Context struct {
	Message string
	Path    string
	Fruits  [3]string
}

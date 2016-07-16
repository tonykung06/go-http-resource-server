package main

import (
	// "fmt"
	"net/http"
	"text/template"
)

func useBasicHtmlTemplate() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		tmpl, err := template.New("test").Parse(doc)
		if err == nil {
			context := Context{"the message"}
			tmpl.Execute(w, context)
		}
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
		<h1>Hello {{.Message}}</h1>
	</body>
</html>
`

type Context struct {
	Message string
}

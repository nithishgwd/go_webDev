package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageVariables struct {
	Title string
	Name  string
}

func main() {
	http.HandleFunc("/", HomePage)
	http.ListenAndServe(":8080", nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		tmpl := template.Must(template.New("index.html").Parse(indexTemplate))
		tmpl.Execute(w, nil)
		return
	}

	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")

	pv := PageVariables{
		Title: "Welcome Page",
		Name:  fmt.Sprintf("%s %s", firstName, lastName),
	}

	tmpl := template.Must(template.New("welcome.html").Parse(welcomeTemplate))
	tmpl.Execute(w, pv)
}

const indexTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>{{.Title}}</title>
</head>
<body>
	<form action="/" method="post">
		<label for="first_name">First Name:</label>
		<input type="text" name="first_name" required>
		<br>
		<label for="last_name">Last Name:</label>
		<input type="text" name="last_name" required>
		<br>
		<input type="submit" value="Submit">
	</form>
</body>
</html>
`

const welcomeTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>{{.Title}}</title>
</head>
<body>
	<h1>Hello, {{.Name}}! Welcome to our website.</h1>
</body>
</html>
`

package web

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageVariables struct {
	Title string
	Name  string
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

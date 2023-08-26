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

	serverName := "Server 1" //server name

    pv := struct {
        PageVariables
        ServerName string
    }{
        PageVariables: PageVariables{
            Title: "Welcome Page",
            Name:  fmt.Sprintf("%s %s", firstName, lastName),
        },
        ServerName: serverName, // Passing the server name to the template
    }

	tmpl := template.Must(template.New("welcome.html").Parse(welcomeTemplate))
	tmpl.Execute(w, pv)
}

func ContactPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handleing the contact form submission
		name := r.FormValue("name")
		email := r.FormValue("email")
		message := r.FormValue("message")

		//save data to DB
		fmt.Printf("Name: %s\nEmail: %s\nMessage: %s\n", name, email, message)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// For GET request, displaying contact page
	tmpl := template.Must(template.New("contact.html").Parse(contactTemplate))
	tmpl.Execute(w, nil)
}

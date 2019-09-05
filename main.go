package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	u := getUserByEmail("bdkinna@gmail.com")
	renderTemplate(w, "index", struct {
		User *user
	}{&u})
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	u := getUserByEmail("bdkinna@gmail.com")
	t := getTasks(u.id)
	renderTemplate(w, "tasks", struct {
		User  *user
		Tasks *[]task
	}{&u, &t})
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tasks/", tasksHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

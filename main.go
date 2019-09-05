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
	t, err := getTasks(u.id)
	if err != nil {
		log.Fatal(err)
	}

	renderTemplate(w, "tasks", struct {
		User  *user
		Tasks *[]task
	}{&u, &t})
}

func taskFormHandler(w http.ResponseWriter, r *http.Request) {
	u := getUserByEmail("bdkinna@gmail.com")
	renderTemplate(w, "task_form", struct {
		User *user
	}{&u})
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	u := getUserByEmail("bdkinna@gmail.com")
	t := task{Name: r.FormValue("name")}

	err := createTask(&u, &t)
	if err != nil {
		log.Fatal(err)
	}

	http.Redirect(w, r, "/tasks/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/tasks/", tasksHandler)
	http.HandleFunc("/tasks/new/", taskFormHandler)
	http.HandleFunc("/tasks/create/", createTaskHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

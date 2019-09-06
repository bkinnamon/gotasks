package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

type templateData struct {
	User  *user
	Task  *task
	Tasks *[]task
}

func renderTemplate(w http.ResponseWriter, tmpl string, data *templateData) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	u := getUserByEmail("bdkinna@gmail.com")
	renderTemplate(w, "index", &templateData{User: &u})
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	u := getUserByEmail("bdkinna@gmail.com")
	t, err := getTasks(u.ID)
	if err != nil {
		log.Fatal(err)
	}

	renderTemplate(w, "tasks", &templateData{User: &u, Tasks: &t})
}

func taskFormHandler(w http.ResponseWriter, r *http.Request) {
	u := getUserByEmail("bdkinna@gmail.com")
	t := &task{Name: ""}
	renderTemplate(w, "task_form", &templateData{User: &u, Task: t})
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	u := getUserByEmail("bdkinna@gmail.com")
	var t *task
	idStr := r.FormValue("id")
	if idStr != "" {
		id, _ := strconv.Atoi(idStr)
		t = getTaskByID(id)
	} else {
		t = &task{Name: r.FormValue("name")}
	}

	err := createTask(&u, t)
	if err != nil {
		log.Fatal(err)
		http.Redirect(w, r, "/tasks/new/", http.StatusFound)
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

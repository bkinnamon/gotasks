package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := getUserByEmail("bdkinna@gmail.com")
	renderTemplate(w, "index", &templateData{User: &u})
}

func tasksHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := getUserByEmail("bdkinna@gmail.com")
	t, err := getTasks(u.ID)
	if err != nil {
		log.Fatal(err)
	}

	renderTemplate(w, "tasks", &templateData{User: &u, Tasks: &t})
}

func taskFormHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := getUserByEmail("bdkinna@gmail.com")
	id, _ := strconv.Atoi(p.ByName("id"))
	t := &task{ID: id, Name: ""}
	renderTemplate(w, "task_form", &templateData{User: &u, Task: t})
}

func createTaskHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := getUserByEmail("bdkinna@gmail.com")
	t := &task{Name: r.FormValue("name")}

	err := createTask(&u, t)
	if err != nil {
		log.Fatal(err)
		http.Redirect(w, r, "/tasks/new/", http.StatusFound)
	}

	http.Redirect(w, r, "/tasks/", http.StatusFound)
}

func main() {
	router := httprouter.New()

	router.GET("/", indexHandler)
	router.GET("/tasks", tasksHandler)
	router.GET("/tasks/new", taskFormHandler)
	router.POST("/tasks", createTaskHandler)
	router.GET("/tasks/:id", taskFormHandler)

	log.Println("Listening ...")
	http.ListenAndServe(":8080", router)
}

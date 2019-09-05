package main

import (
	"fmt"
	"html/template"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	// printUsers()

	u := getUserByEmail("bdkinna@gmail.com")

	fmt.Printf("[%s] %s: %s", u.id, u.name, u.email)
}

package main

import (
	"fmt"
	"log"
)

type user struct {
	id    string
	Email string
	Name  string
}

func printUsers() {
	initDb()

	sql := "SELECT id, email, name FROM users;"
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var u user

		if err := rows.Scan(&u.id, &u.Email, &u.Name); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[%s] %s: %s\n", u.id, u.Name, u.Email)
	}
}

func getUserByEmail(email string) user {
	initDb()

	sql := "SELECT id, email, name FROM users WHERE email = $1"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	var u user

	err = stmt.QueryRow(email).Scan(&u.id, &u.Email, &u.Name)
	if err != nil {
		log.Fatal(err)
	}

	return u
}

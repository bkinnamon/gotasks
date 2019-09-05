package main

import (
	"log"
)

type task struct {
	id         int
	Name       string
	IsComplete bool
}

func getTasks(userID int) ([]task, error) {
	initDb()

	sql := "SELECT id, name, is_complete FROM tasks WHERE user_id = $1"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []task

	rows, err := stmt.Query(userID)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var t task
		err := rows.Scan(&t.id, &t.Name, &t.IsComplete)
		if err != nil {
			log.Fatal(err)
		}

		tasks = append(tasks, t)
	}

	return tasks, err
}

func createTask(u *user, t *task) error {
	initDb()

	sql := "INSERT INTO tasks (name, user_id) VALUES ($1, $2)"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(t.Name, u.id)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

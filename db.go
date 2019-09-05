package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDb() {
	if db == nil {
		var connStr = "user=postgres password=docker dbname=gotasks sslmode=disable"

		database, err := sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}

		db = database
	}
}

func fetchRow(columns string, table string, where string) *sql.Row {
	initDb()

	sql := "SELECT " + columns + " FROM " + table + " WHERE " + where + ";"

	row := db.QueryRow(sql)

	return row
}

func fetchRows(columns string, table string) *sql.Rows {
	initDb()

	rows, err := db.Query("SELECT " + columns + " FROM " + table + ";")
	if err != nil {
		panic(err)
	}

	return rows
}

func testDb() {
	initDb()

	rows, err := db.Query("SELECT name FROM users;")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

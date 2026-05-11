package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "gee.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err := db.Exec("DROP TABLE IF EXISTS User;"); err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec("CREATE TABLE User(Name text);"); err != nil {
		log.Fatal(err)
	}
	result, err := db.Exec("INSERT INTO User(`Name`) values (?), (?)", "TOM", "Sam")

	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}

	row := db.QueryRow("SELECT Name FROM User LIMIT 1")
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println(name)
	}
}

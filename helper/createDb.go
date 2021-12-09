package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "main.db")

	if err != nil {
		panic(err)
	}

	currentDir, _ := os.Getwd()

	file, err := os.ReadFile(currentDir + "/adapter/repository/fixture/sql/1-transactions.up.sql")

	if err != nil {
		panic(err)
	}

	db.Exec(string(file))

	db.Close()
}

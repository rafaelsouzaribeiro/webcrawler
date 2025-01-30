package repository

import (
	"log"
)

func (q *Isqlite) CreateTable() {
	createTableSQL := `CREATE TABLE IF NOT EXISTS visited_pages (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
        "url" TEXT NOT NULL
    );`

	_, err := q.Db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

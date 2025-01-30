package repository

import (
	"log"
)

func (q *Isqlite) InsertVisitedPage(url string) {
	insertPageSQL := `INSERT INTO visited_pages(url) VALUES (?)`
	statement, err := q.Db.Prepare(insertPageSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	_, err = statement.Exec(url)
	if err != nil {
		log.Fatal(err)
	}
}

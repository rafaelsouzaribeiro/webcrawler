package sqlite

import (
	"database/sql"
	"log"
)

func InsertVisitedPage(db *sql.DB, url string) {
	insertPageSQL := `INSERT INTO visited_pages(url) VALUES (?)`
	statement, err := db.Prepare(insertPageSQL)
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	_, err = statement.Exec(url)
	if err != nil {
		log.Fatal(err)
	}
}

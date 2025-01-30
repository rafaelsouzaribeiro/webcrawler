package sqlite

import (
	"database/sql"
	"log"
)

func IsPageVisited(db *sql.DB, url string) bool {
	query := `SELECT COUNT(*) FROM visited_pages WHERE url = ?`
	var count int
	err := db.QueryRow(query, url).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

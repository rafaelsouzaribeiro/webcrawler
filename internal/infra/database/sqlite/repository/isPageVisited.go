package repository

import (
	"log"
)

func (q *Isqlite) IsPageVisited(url string) bool {
	query := `SELECT COUNT(*) FROM visited_pages WHERE url = ?`
	var count int
	err := q.Db.QueryRow(query, url).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count > 0
}

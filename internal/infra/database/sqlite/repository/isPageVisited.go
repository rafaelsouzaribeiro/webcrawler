package repository

func (q *Isqlite) IsPageVisited(url string) (bool, error) {
	query := `SELECT COUNT(*) FROM visited_pages WHERE url = ?`
	var count int
	err := q.Db.QueryRow(query, url).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

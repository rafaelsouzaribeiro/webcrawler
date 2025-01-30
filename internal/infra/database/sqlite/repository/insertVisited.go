package repository

func (q *Isqlite) InsertVisitedPage(url string) error {
	insertPageSQL := `INSERT INTO visited_pages(url) VALUES (?)`
	statement, err := q.Db.Prepare(insertPageSQL)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(url)
	if err != nil {
		return err
	}

	return nil
}

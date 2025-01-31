package connection

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

func GetSqliteDataBase(filepath string) (*sql.DB, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		file, err := os.Create(filepath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

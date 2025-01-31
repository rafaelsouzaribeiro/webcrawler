package connection

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func GetSqliteDataBase(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", filepath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

package connection

import (
	"database/sql"
)

func GetSqliteDataBase(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

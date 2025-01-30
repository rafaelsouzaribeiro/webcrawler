package connection

import (
	"database/sql"
	"log"
)

func GetSqliteDataBase(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

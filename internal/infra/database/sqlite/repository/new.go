package repository

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

type Isqlite struct {
	Db *sql.DB
}

func NewSqliteRepository(c *sql.DB) *Isqlite {
	return &Isqlite{
		Db: c,
	}
}

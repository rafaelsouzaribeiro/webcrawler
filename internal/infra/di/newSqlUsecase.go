package di

import (
	"database/sql"

	"github.com/rafaelsouzaribeiro/webcrawler/internal/infra/database/sqlite/repository"
	"github.com/rafaelsouzaribeiro/webcrawler/internal/usecase"
)

func NewSqlUseCase(db *sql.DB) usecase.SqlUsecase {
	repo := repository.NewSqliteRepository(db)
	first := usecase.NewSqlUsecase(repo)
	return *first
}

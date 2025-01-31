package usecase

import "github.com/rafaelsouzaribeiro/webcrawler/internal/ports/irepository"

type SqlUsecase struct {
	Repository irepository.IRepository
}

func NewSqlUsecase(repository irepository.IRepository) *SqlUsecase {
	return &SqlUsecase{Repository: repository}
}

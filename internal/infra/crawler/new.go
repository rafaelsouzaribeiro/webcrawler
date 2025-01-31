package crawler

import (
	"github.com/rafaelsouzaribeiro/webcrawler/internal/usecase"
)

type Usecasse struct {
	usecase usecase.SqlUsecase
}

func NewCrawler(
	repos usecase.SqlUsecase,
) *Usecasse {
	return &Usecasse{
		usecase: repos,
	}
}

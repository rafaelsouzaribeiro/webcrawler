package di

import (
	"github.com/rafaelsouzaribeiro/webcrawler/internal/infra/database/redis/repository"
	"github.com/rafaelsouzaribeiro/webcrawler/internal/usecase"
	"github.com/redis/go-redis/v9"
)

func NewRedisUseCase(db *redis.Client) usecase.SqlUsecase {
	repo := repository.NewRedisRepository(db)
	first := usecase.NewSqlUsecase(repo)
	return *first
}

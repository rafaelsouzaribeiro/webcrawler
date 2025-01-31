package main

import (
	"github.com/rafaelsouzaribeiro/webcrawler/internal/infra/crawler"
	"github.com/rafaelsouzaribeiro/webcrawler/internal/infra/database/sqlite/connection"
	"github.com/rafaelsouzaribeiro/webcrawler/internal/infra/di"
)

func main() {
	receive := make(chan string)

	db, err := connection.GetSqliteDataBase("database.db")

	if err != nil {
		panic(err)
	}

	usecase := di.NewSqlUseCase(db)
	usecase.Repository.CreateTable()
	execute := crawler.NewCrawler(usecase)
	execute.InitCrawler("https://ge.globo.com/", receive)

	for element := range receive {
		println(element)
	}

}

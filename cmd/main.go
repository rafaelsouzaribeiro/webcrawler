package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
	"github.com/rafaelsouzaribeiro/webcrawler/internal/infra/crawler"
	"github.com/rafaelsouzaribeiro/webcrawler/internal/infra/database/sqlite/connection"
	"github.com/rafaelsouzaribeiro/webcrawler/internal/infra/di"
	"github.com/rafaelsouzaribeiro/webcrawler/pkg/redis/consumer"
)

func main() {
	receive := make(chan string)
	channelConsumer := make(chan payload.Message)

	db, err := connection.GetSqliteDataBase("database.db")

	if err != nil {
		panic(err)
	}

	usecase := di.NewSqlUseCase(db)
	execute := crawler.NewCrawler(usecase)
	execute.InitCrawler("https://ge.globo.com/", receive)

	consumer.ConsumerRedis(&[]string{"crawler"}, "localhost:6379", channelConsumer)

	for msgs := range channelConsumer {
		execute.InitCrawler(string(msgs.Value), receive)
		fmt.Printf("Message: %s Topic:%s\n", string(msgs.Value), msgs.Topic)
	}

	close(channelConsumer)

	select {}

}

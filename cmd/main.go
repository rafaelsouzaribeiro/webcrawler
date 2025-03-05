package main

import (
	"fmt"

	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
	"github.com/rafaelsouzaribeiro/webcrawler/internal/infra/crawler"
	"github.com/rafaelsouzaribeiro/webcrawler/internal/infra/database/redis/connection"
	"github.com/rafaelsouzaribeiro/webcrawler/internal/infra/di"
	"github.com/rafaelsouzaribeiro/webcrawler/pkg/redis/consumer"
)

func main() {
	receive := make(chan string)
	channelConsumer := make(chan payload.Message)

	// db, err := connection.GetSqliteDataBase("database.db")

	// if err != nil {
	// 	panic(err)
	// }

	db := connection.ConnectingRedis("localhost", 6379, "123mudar")

	//usecase := di.NewSqlUseCase(db)
	usecase := di.NewRedisUseCase(db)
	execute := crawler.NewCrawler(usecase)
	execute.InitCrawler("https://ge.globo.com/", receive)

	consumer.ConsumerRedis(&[]string{"crawler_broker"}, "localhost:6379", channelConsumer)

	for msgs := range channelConsumer {
		execute.InitCrawler(string(msgs.Value), receive)
		fmt.Printf("Message: %s Topic:%s\n", string(msgs.Value), msgs.Topic)
	}

	close(channelConsumer)

}

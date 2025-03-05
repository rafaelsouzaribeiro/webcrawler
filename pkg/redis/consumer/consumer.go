package consumer

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func ConsumerRedis(topics *[]string, host string, channelConsumer chan<- payload.Message) {
	data := payload.Message{
		Topics: topics,
	}

	broker := factory.NewBroker(factory.Redis, host)
	go broker.Consumer(&data, channelConsumer)

}

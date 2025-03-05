package producer

import (
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/factory"
	"github.com/rafaelsouzaribeiro/golang-broker/pkg/payload"
)

func Producer(value, topic, host string) {

	message := payload.Message{
		Value: []byte(value),
		Topic: topic,
		Headers: &[]payload.Header{
			{
				Key:   "your-header-key1",
				Value: "your-header-value1",
			},
			{
				Key:   "your-header-key2",
				Value: "your-header-value2",
			},
		},
	}

	pro := factory.NewBroker(factory.Redis, host)
	pro.SendMessage(&message)

}

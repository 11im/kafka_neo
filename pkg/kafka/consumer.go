package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

func CreateConsumer() sarama.Consumer {
	log.Println("Create Consumer Instance")
	con, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)

	if err != nil {
		panic(err)
	}

	return con
}

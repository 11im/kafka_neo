package kafka

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
	util "github.com/ijh4565/kafka_neo/pkg/util"
)

func ConsumePartitionLambda(topic string) {
	type info util.Info
	con := KafkaConsumer()
	log.Println("Start Consuming")
	pCon, err := con.ConsumePartition(topic, 0, sarama.OffsetOldest)

	if err != nil {
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

ConsumerLoop:
	for {
		select {
		case msg := <-pCon.Messages():
			info := util.JsonConvert(msg.Value)

			log.Println(info)
		case <-signals:
			break ConsumerLoop
		}
	}
}

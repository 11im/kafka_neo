package kafka

import (
	neo "github.com/ijh4565/kafka_neo/pkg/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
	util "github.com/ijh4565/kafka_neo/pkg/util"
)

func ConsumePartitionLambda(topic string, client neo4j.Driver, wg sync.WaitGroup) {
	type info util.Info
	con := CreateConsumer()
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
			neo.Neo4jWriteLambda(client, info)
			log.Println(info)
		case <-signals:
			wg.Done()
			break ConsumerLoop
		}
	}
}

package kafka

import (
	"github.com/Shopify/sarama"
	neo "github.com/ijh4565/kafka_neo/pkg/neo4j"
	util "github.com/ijh4565/kafka_neo/pkg/util"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
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
	defer wg.Done()

	var totalProcessingTime int64 = 0

ConsumerLoop:
	for {
		select {
		case msg := <-pCon.Messages():
			info := util.JsonConvert(msg.Value)
			start := time.Now().UnixMicro()
			neo.Neo4jWriteLambda(client, info)
			end := time.Now().UnixMicro() - start
			totalProcessingTime = totalProcessingTime + end
			log.Println("Lambda Insert Accumulate Processing Time :", totalProcessingTime)

		case <-signals:
			break ConsumerLoop
		}
	}
}

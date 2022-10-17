package kafka

import (
	neo "github.com/ijh4565/kafka_neo/pkg/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"io"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	util "github.com/ijh4565/kafka_neo/pkg/util"
)

func ConsumePartitionKappa(topic, insertLogPath string, client neo4j.Driver, wg sync.WaitGroup) {
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

	logFile, err := os.OpenFile(insertLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiLogWriter := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multiLogWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
ConsumerLoop:
	for {
		select {
		case msg := <-pCon.Messages():
			info := util.JsonConvert(msg.Value)
			start := time.Now().UnixMicro()
			neo.Neo4jWriteKappa(client, info)
			end := time.Now().UnixMicro() - start
			totalProcessingTime = totalProcessingTime + end
			log.Println("Kappa Insert Accumulate Processing Time :", totalProcessingTime)

		case <-signals:
			break ConsumerLoop
		}
	}
}

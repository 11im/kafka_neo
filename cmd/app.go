package main

import (
	"github.com/ijh4565/kafka_neo/pkg/kafka"
	neo "github.com/ijh4565/kafka_neo/pkg/neo4j"
	"sync"
)

func Kappa(insertLogPath string) {
	var wg sync.WaitGroup
	client1 := neo.Neo4JClient("neo4j", "neo4j1")
	client2 := neo.Neo4JClient("neo4j", "neo4j1")

	wg.Add(1)
	go kafka.ConsumePartitionKappa("neo_connect", insertLogPath, client1, wg)
	go neo.Neo4jQuery(client2)
	wg.Wait()
	defer client1.Close()
	defer client2.Close()
}

func Lambda(insertLogPath, batchLogPath string) {
	var wg sync.WaitGroup
	wg.Add(1)
	client1 := neo.Neo4JClient("neo4j", "neo4j1")
	client2 := neo.Neo4JClient("neo4j", "neo4j1")
	go kafka.ConsumePartitionLambda("neo_connect", insertLogPath, client1, wg)
	go neo.Neo4JLambdaBatch(client1, batchLogPath)
	go neo.Neo4jQuery(client2)
	wg.Wait()
	defer client1.Close()
	defer client2.Close()
}

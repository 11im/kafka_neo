package main

import (
	"github.com/ijh4565/kafka_neo/pkg/kafka"
	neo "github.com/ijh4565/kafka_neo/pkg/neo4j"
	"sync"
)

func Kappa() {
	var wg sync.WaitGroup
	client := neo.Neo4JClient("neo4j", "neo4j1")
	wg.Add(1)
	go kafka.ConsumePartitionKappa("neo_connect", client, wg)
	wg.Wait()
	defer client.Close()
}

func Lambda() {
	var wg sync.WaitGroup
	wg.Add(1)
	client := neo.Neo4JClient("neo4j", "neo4j1")
	go kafka.ConsumePartitionLambda("neo_connect", client, wg)
	go neo.Neo4JLambdaBatch(client)

	wg.Wait()
	defer client.Close()
}

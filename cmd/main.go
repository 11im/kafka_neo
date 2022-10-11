package main

import(
	"fmt"
	"os"
	kafka "https://github.com/ijh4565/kafka_neo/pkg/kafka"
	neo "https://github.com/ijh4565/kafka_neo/pkg/neo4j"
)


func main(){
	flag := os.Args[1]

	if flag == "K"{
		kafka.Kappa("neo_connect")
	} else if flag == "L"{
		kafka.Lambda("neo_connect")
	}
}
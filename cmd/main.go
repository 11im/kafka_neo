package main

import (
	"fmt"
	"os"

	kafka "github.com/ijh4565/kafka_neo/pkg/kafka"
)

func main() {
	flag := os.Args[1]

	if flag == "K" {
		kafka.Kappa("neo_connect")
	} else if flag == "L" {
		kafka.Lambda("neo_connect")
	} else {
		fmt.Println("----------")
		fmt.Println("-Use with Flag-")
		fmt.Println("-K : kappa-")
		fmt.Println("-L : lambda-")
		fmt.Println("----------")
	}
}

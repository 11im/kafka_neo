package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("----------")
		fmt.Println("-Use with Flag-")
		fmt.Println("-K : kappa-")
		fmt.Println("-L : lambda-")
		fmt.Println("----------")
		panic("Not Enough Arguments")
	}

	flag := os.Args[1]
	if flag == "K" {
		logPath := os.Args[2]
		Kappa(logPath)
	} else if flag == "L" {
		insertLogPath := os.Args[2]
		batchLogPath := os.Args[3]
		Lambda(insertLogPath, batchLogPath)
	} else {
		fmt.Println("----------")
		fmt.Println("-Use with Flag, Log Path-")
		fmt.Println("-K : kappa, Insert Log Path-")
		fmt.Println("-L : lambda, Insert Log Path, Batch Log Path-")
		fmt.Println("----------")
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	flag := os.Args[1]

	if flag == "K" {
		Kappa()
	} else if flag == "L" {
		Lambda()
	} else {
		fmt.Println("----------")
		fmt.Println("-Use with Flag-")
		fmt.Println("-K : kappa-")
		fmt.Println("-L : lambda-")
		fmt.Println("----------")
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	flag := os.Args[1]
	if len(os.Args) < 2 {
		fmt.Println("----------")
		fmt.Println("-Use with Flag-")
		fmt.Println("-K : kappa-")
		fmt.Println("-L : lambda-")
		fmt.Println("----------")
		panic("Not Enough Arguments")
	}
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

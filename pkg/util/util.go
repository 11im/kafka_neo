package util

import (
	"encoding/json"
)

type Node struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Edge struct {
	Weight float32 `json:"weight"`
}

type Info struct {
	Id    string `json:"id"`
	Node1 Node
	Node2 Node
	Edge  Edge
}

func JsonConvert(input []byte) Info {
	var info Info
	json.Unmarshal(input, &info)
	return info
}

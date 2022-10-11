package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

// Neo4J client
func Neo4JClient(username, password string) neo4j.Driver {
	dbUri := "neo4j://localhost:7687"
	driver, err := neo4j.NewDriver(dbUri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		panic(err)
	}
	return driver
}

package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
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

func Neo4jQuery(driver neo4j.Driver) {

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.Run("$Query", map[string]interface{}{
		"nid1":   info.Node1.Id,
		"name1":  info.Node1.Name,
		"nid2":   info.Node2.Id,
		"name2":  info.Node2.Name,
		"eid1":   info.Edge.Id,
		"weight": info.Edge.Weight,
	})
	if err != nil {
		panic(err)
	}
	log.Println(result)
}

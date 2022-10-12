package neo4j

import (
	util "github.com/ijh4565/kafka_neo/pkg/util"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
)

func Neo4jWriteKappa(driver neo4j.Driver, info util.Info) {

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result, err := session.Run("MERGE (a:Person {id: $nid1, name: $name1}) MERGE (b:Person {id: $nid2, name: $name2}) CREATE (a)-[r:follow {id: $eid1, weight: $weight}]->(b)", map[string]interface{}{
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

package neo4j

import (
	"fmt"
	util "github.com/ijh4565/kafka_neo/pkg/util"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"time"
)

func Neo4jWriteLambda(driver neo4j.Driver, info util.Info) {

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result, err := session.Run("CREATE p = (a:Person {id: $nid1, name: $name1})-[r: follow {id: $eid1, weight: $weight}]->(b: Person {id:$nid2, name: $name2})", map[string]interface{}{
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

func Neo4JLambdaBatch(driver neo4j.Driver) {

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	for true {
		time.Sleep(time.Second * 10)
		result, err := session.Run("MATCH (n: Person) WITH n.id AS id, COLLECT(n) AS nodelist, COUNT(*) AS count WHERE count > 1 CALL apoc.refactor.mergeNodes(nodelist) YIELD node RETURN count(node)", map[string]interface{}{})
		if err != nil {
			panic(err)
		}
		fmt.Println("Batch : ", result)
	}
}

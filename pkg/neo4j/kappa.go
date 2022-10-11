package neo4j

import (
	util "https://github.com/ijh4565/kafka_neo/pkg/util"
)

func CreateItemKappa(tx neo4j.Transaction, con util.Info) error {
	// Query 들어갈 자리

	if err != nil {
		return err
	}
	return nil
}

func Neo4jWriteKappa(driver neo4j.Driver, info util.Info) {

	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result, err := session.RunRun("CREATE (n:Node {id: $id, name: $name })", map[string]interface{}{
		"id":   con.Node1.Id,
		"name": con.Node1.Name,
	})
	if err != nil {
		panic(err)
	}
}

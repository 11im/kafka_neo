package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	util "https://github.com/ijh4565/kafka_neo/util"
)

func CreateItemLamdba(tx neo4j.Transaction, con util.Info) error {
	records, err := tx.Run("CREATE (n:Node {id: $id, name: $name })", map[string]interface{}{
		"id":   con.Node1.Id,
		"name": con.Node1.Name,
	})
	// Query 들어갈 자리

	if err != nil {
		return err
	}
	return nil
}

func Neo4jWriteLambda(driver neo4j.Driver) {

	defer driver.Close()

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()

	result, err := session.WriteTransaction(CreateItemKappa)
	if err != nil {
		panic(err)
	}
}

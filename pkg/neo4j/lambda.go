package neo4j

import (
	util "github.com/ijh4565/kafka_neo/pkg/util"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"io"
	"log"
	"os"
	"time"
)

func Neo4jWriteLambda(driver neo4j.Driver, info util.Info) {

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	_, err := session.Run("CREATE p = (a:Person {id: $nid1, name: $name1})-[r: follow {id: $eid1, weight: $weight}]->(b: Person {id:$nid2, name: $name2})", map[string]interface{}{
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
}

func Neo4JLambdaBatch(driver neo4j.Driver, batchLogPath string) {

	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	var totalProcessingTime int64 = 0

	logFile, err := os.OpenFile(batchLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	multiLogWriter := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multiLogWriter)

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	for true {
		time.Sleep(time.Second * 10)
		start := time.Now().UnixMicro()
		_, err := session.Run("MATCH (n: Person) WITH n.id AS id, COLLECT(n) AS nodelist, COUNT(*) AS count WHERE count > 1 CALL apoc.refactor.mergeNodes(nodelist) YIELD node RETURN count(node)", map[string]interface{}{})
		end := time.Now().UnixMicro() - start
		totalProcessingTime = totalProcessingTime + end
		if err != nil {
			panic(err)
		}
		log.Println("Total Lambda Batch Processing Time : ", totalProcessingTime)
	}
}

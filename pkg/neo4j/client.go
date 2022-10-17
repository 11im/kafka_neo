package neo4j

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"math/rand"
	"time"
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

	for true {
		time.Sleep(time.Second)
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

		randStartName := chars[r1.Intn(len(chars))]
		randEndName := chars[r1.Intn(len(chars))]
		randRangeName := chars[r1.Intn(len(chars))]
		randWeight := float64(r1.Intn(5)) + r1.Float64()

		result, err := session.Run("MATCH (n: Person)-[r: follow]->(m: Person) WHERE r.weight < $weight RETURN n.name, m.name", map[string]interface{}{
			"weight": randWeight,
		})
		if err != nil {
			panic(err)
		}
		log.Println("Q1 ", result.Record())

		result, err = session.Run("MATCH (n: Person) WHERE n.name STARTS WITH $name RETURN n.name", map[string]interface{}{
			"name": randStartName,
		})
		if err != nil {
			panic(err)
		}
		log.Println("Q2 ", result.Record())

		result, err = session.Run("MATCH (n: Person) WHERE n.name ENDS WITH $name RETURN n.name", map[string]interface{}{
			"name": randEndName,
		})
		if err != nil {
			panic(err)
		}
		log.Println("Q3 ", result.Record())

		result, err = session.Run("MATCH (n: Person) WHERE n.name >= $name RETURN n.name", map[string]interface{}{
			"name": randRangeName,
		})
		if err != nil {
			panic(err)
		}
		log.Println("Q4 ", result.Record())

		result, err = session.Run("MATCH (n: Person)-[r: follow]->(m: Person) RETURN count(n), count(r)", map[string]interface{}{
		})
		if err != nil {
			panic(err)
		}
		log.Println("Q5 ", result.Record())
	}
}



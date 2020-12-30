package main

import (
	"fmt"
	"github.com/nargetdev/neostate/neo"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func main() {
	n := neo.Neo{}
	n.ConnectToDB()

	label := "CONNECTED_TO"
	myCmd := `//MATCH(a:EVENT) WHERE a.name=$EventName
	CREATE (n:INCHARGE {name:$name, gender: $gender})<-[:`+label+`]-(a:EVENT {name: "Ahau"})`

	result, err := n.Session.Run(myCmd, map[string]interface{}{
		"EventName":   "Ahau",
		"name": "Item 21",
		"gender": "Whatsthat",
	})
	if err != nil {
		fmt.Println(err)
	}


	var record *neo4j.Record
	for result.NextRecord(&record) {
		fmt.Printf("Created Item with Id = '%d' and Name = '%s'\n", record.Values[0].(int64), record.Values[1].(string))
	}
}
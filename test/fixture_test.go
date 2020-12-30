package test

import (
	"fmt"
	"github.com/SynaesthesiaFlow/stateful-mesh-service/pkg/fixture"
	"github.com/nargetdev/neostate/neo"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"testing"
)

func TestFixture(t *testing.T) {
	f := fixture.CreateStrip(.5, 10)

	for p := range f.Points {
		//graphene.NeoPoint(p)
		fmt.Println(p)
	}



	label := "CONNECTED_TO"
	myCmd := `
	MATCH(a:EVENT) WHERE a.name=$EventName
	CREATE (n:INCHARGE {name:$name, gender: $gender})<-[:`+label+`]-(a)
	RETURN n, a, id(n), id(a)
	`
	myVars := map[string]interface{}{
		"EventName":   "Ahau",
		"name": "Item 21",
		"gender": "Whatsthat",
	}

	n := neo.Neo{}
	n.ConnectToDB()
	result, err := n.Session.Run(myCmd, myVars)
	if err != nil {
		fmt.Println(err)
	}


	var record *neo4j.Record
	for result.NextRecord(&record) {
		fmt.Printf("Created Item with Id = '%v'\n", record.Values)
	}

}
package test

import (
	"fmt"
	"github.com/SynaesthesiaFlow/stateful-mesh-service/pkg/fixture"
	"github.com/nargetdev/neostate/graphene"
	"github.com/nargetdev/neostate/neo"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"testing"
)

func TestFixture(t *testing.T) {
	f := fixture.CreateStrip(.5, 10)





	n := neo.Neo{}
	n.ConnectToDB()

	//label := "CONNECTED_TO"
	myCmd := ""
	myVars := map[string]interface{}{}

	lastNodeId := int64(-1)

	for _, p := range f.Points {
		myCmd = graphene.NeoPoint(p)

		result, err := n.Session.Run(myCmd, myVars)
		if err != nil {
			fmt.Println(err)
		}


		//if lastNodeId >= 0 { // -1 if unset
		//	myCmd =
		//}
		//
		thisNodeId := int64(-1)
		var record *neo4j.Record
		for result.NextRecord(&record) {
			node, ok :=record.Values[0].(dbtype.Node)
			if ok {
				thisNodeId = node.Id
			}
			fmt.Println(lastNodeId)
		}

		// OK Connect the dots...
		myCmd = graphene.NeoRelationship("FOLLOWS", int(thisNodeId), int(lastNodeId))
		fmt.Println(myCmd)
		result, err = n.Session.Run(myCmd, myVars)
		if err != nil {
			fmt.Println(err)
		}

		lastNodeId = thisNodeId
	}
}
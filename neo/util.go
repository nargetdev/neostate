package neo

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

/**
 * Common utility functions for building up neo structures
 *
 */

func (n Neo) CreateNode (s string) (id uint32, err error) {
	myCmd := `//MATCH(a:EVENT) WHERE a.name=$EventName
	CREATE (n:INCHARGE {name:$name})`
	myVars := map[string]interface{}{
		"EventName":   "Ahau",
		"name": "Item 21",
		"gender": "Whatsthat",
	}

	result, err := n.Session.Run(myCmd, myVars)
	if err != nil {
		fmt.Println(err)
	}

	var record *neo4j.Record
	for result.NextRecord(&record) {
		fmt.Printf("Created Item with Id = '%d' and Name = '%s'\n", record.Values[0].(int64), record.Values[1].(string))
	}
	return 0xffff, nil
}

func GenerateCmd(objIn interface{}) (s string, err error) {
	switch v := objIn.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n%v", v, v)
	}
	return "", nil
}
package graphene

import (
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)


func NeoPoint (p neo4j.Point3D) string {
	resultString := ""

	startStr :=  "CREATE (p :Point "
	resultString += startStr

	params := generatePointParams(p)
	//params := "{point: [0, 0, 0]}"
	resultString += params

	closeStr := ") RETURN p"
	resultString += closeStr

	//fmt.Println(resultString)
	return resultString
}

func generatePointParams(p neo4j.Point3D) string {
	resultString := ""
	resultString += "{location: ["
	resultString += fmt.Sprintf("%f", float64(p.X)) + ","
	resultString += fmt.Sprintf("%f", float64(p.Y)) + ","
	resultString += fmt.Sprintf("%f", float64(p.Z))
	resultString += "]}"

	return resultString
}

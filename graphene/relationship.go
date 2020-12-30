package graphene

import "strconv"

func NeoRelationship (label string, id0 int, id1 int) string {
	return `MATCH (a) MATCH (b)
WHERE id(a)=` + strconv.Itoa(id0) + ` and id(b)=`+ strconv.Itoa(id1) + `
CREATE (a)-[:` + label + `]->(b)
RETURN a, b`
}

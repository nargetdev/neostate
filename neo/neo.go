package neo

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type Neo struct {
	user    string
	pass    string
	url     string
	Session neo4j.Session
	driver  neo4j.Driver
}
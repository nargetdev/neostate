package graphene

func ConnectionById (id int) string {
	matchString := `
	MATCH (n)
	WHERE id(n)=` + string(rune(id)) + `
	`
	nodeToCreate := ""
	connectionToCreate := ""

	createString := `CREATE (n)` + connectionToCreate + nodeToCreate

	resultString := matchString + createString
	return resultString

}
package neo

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

func (n *Neo) ConnectToDB() (neo4j.Session, neo4j.Driver, error) {
	// define driver, Session and result vars
	var (
		driver  neo4j.Driver
		session neo4j.Session
		err     error
	)
	// initialize driver to connect to localhost with ID and password
	if driver, err = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "password", "")); err != nil {
		return nil, nil, err
	}
	// Open a new Session with write access
	session = driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	n.Session = session
	n.driver = driver

	return session, driver, nil
}


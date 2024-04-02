package cassandra

import "github.com/monzo/gocassa"

// getKeySpace creates a connection to the Cassandra DB instance
func GetKeySpace() (gocassa.KeySpace, error) {
	keySpace, err := gocassa.ConnectToKeySpace("logs", []string{"127.0.0.1"}, "", "")
	if err != nil {
		panic(err)
	}

	return keySpace, nil
}

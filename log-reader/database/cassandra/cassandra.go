package cassandra

import (
	"github.com/mhope-2/log-reader/shared"
	"github.com/monzo/gocassa"
)

// GetLogsKeySpace creates a connection to the Cassandra DB instance
func GetLogsKeySpace() (gocassa.KeySpace, error) {
	env := shared.NewEnvEnvConfig()

	keySpace, err := gocassa.ConnectToKeySpace("logs", []string{env.CassandraDBHost}, "", "")
	if err != nil {
		panic(err)
	}

	return keySpace, nil
}

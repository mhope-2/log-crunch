package cassandra

import (
	"github.com/mhope-2/log-query/shared"
	"github.com/monzo/gocassa"
)

// getKeySpace creates a connection to the Cassandra DB instance
func GetLogsKeySpace() gocassa.KeySpace {
	env := shared.NewEnvEnvConfig()

	keySpace, err := gocassa.ConnectToKeySpace("logs", []string{env.CassandraDBHost}, "", "")
	if err != nil {
		panic(err)
	}

	return keySpace
}

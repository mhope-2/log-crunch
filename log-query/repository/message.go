package repository

import (
	"context"
	"github.com/mhope-2/log-query/database/cassandra"
	"github.com/mhope-2/log-query/types"
	"github.com/monzo/gocassa"
	"log"
)

func RetrieveMessage(id string) types.Message {
	result := types.Message{}
	keySpace, err := cassandra.GetLogsKeySpace()
	if err != nil {
		panic(err)
	}

	salesTable := keySpace.Table("message", &types.Message{}, gocassa.Keys{
		PartitionKeys: []string{"ID"},
	})

	if err := salesTable.Where(gocassa.Eq("ID", id)).ReadOne(&result).RunWithContext(context.TODO()); err != nil {
		log.Printf("Error querying data: %v", err)
	}

	return result
}

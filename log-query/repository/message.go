package repository

import (
	"context"
	"github.com/mhope-2/log-query/database/cassandra"
	"github.com/mhope-2/log-query/types"
	"github.com/monzo/gocassa"
	"log"
)

func (r *Repository) RetrieveMessage(id string) (types.Message, error) {
	result := types.Message{}
	keySpace, err := cassandra.GetLogsKeySpace()
	if err != nil {
		log.Fatalf("Failed to connect to logs KeySpace: %v", err)
		return types.Message{}, err
	}

	salesTable := keySpace.Table("messages", &types.Message{}, gocassa.Keys{
		PartitionKeys: []string{"ID"},
	})

	if err := salesTable.Where(gocassa.Eq("ID", id)).ReadOne(&result).RunWithContext(context.TODO()); err != nil {
		log.Printf("Error querying data: %v", err)
		return types.Message{}, err
	}

	return result, nil
}

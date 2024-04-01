package cassandra

import (
	"context"
	"github.com/google/uuid"
	"github.com/mhope-2/log-reader/database/types"
	"github.com/monzo/gocassa"
	"log"
)

// getKeySpace creates a connection to the Cassandra DB instance
func getKeySpace() (gocassa.KeySpace, error) {
	keySpace, err := gocassa.ConnectToKeySpace("logs", []string{"127.0.0.1"}, "", "")
	if err != nil {
		panic(err)
	}

	return keySpace, nil
}

func saveToMessageTable(msg *types.Message) {
	keySpace, err := getKeySpace()

	salesTable := keySpace.Table("message", &types.Message{}, gocassa.Keys{
		PartitionKeys: []string{"ID"},
	})

	err = salesTable.Set(types.Message{
		ID:        uuid.New().String(),
		Level:     msg.Level,
		Value:     msg.Value,
		Timestamp: msg.Timestamp,
	}).RunWithContext(context.TODO())
	if err != nil {
		panic(err)
	}

	result := types.Message{}
	if err := salesTable.Where(gocassa.Eq("value", msg.Value)).ReadOne(&result).RunWithContext(context.TODO()); err != nil {
		panic(err)
	}
	log.Println(result)
}

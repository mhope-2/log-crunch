package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/mhope-2/log-reader/database/cassandra"
	"github.com/mhope-2/log-reader/database/types"
	"github.com/monzo/gocassa"
	"log"
)

// SaveMessage inserts a message record into the messages table
func SaveMessage(msg *types.Message) error {
	keySpace, err := cassandra.GetLogsKeySpace()

	messagesTable := keySpace.Table("messages", &types.Message{}, gocassa.Keys{
		PartitionKeys: []string{"id"},
	})

	err = messagesTable.CreateIfNotExist()
	if err != nil {
		log.Fatalf("Failed to create messages table: %v", err)
		return err
	}

	err = messagesTable.Set(types.Message{
		ID:        uuid.New().String(),
		LogLevel:  msg.LogLevel,
		Value:     msg.Value,
		Timestamp: msg.Timestamp,
	}).RunWithContext(context.TODO())
	if err != nil {
		log.Printf("Error saving mesage to db; err: %v", err)
		return err
	}

	return err
}

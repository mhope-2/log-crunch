package kafka

import (
	"context"
	"fmt"
	"github.com/mhope-2/log-reader/shared"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type Consumer struct {
	reader *kafka.Reader
}

func NewConsumer() *Consumer {
	env := shared.NewEnvEnvConfig()

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     env.KafkaBrokers,
		Topic:       env.Topic,
		GroupID:     env.GroupID,
		StartOffset: kafka.FirstOffset, // from whence the consumer group should begin consuming
	})

	return &Consumer{
		reader: r,
	}
}

func (c *Consumer) Consume(ctx context.Context) error {
	for {
		// read message from the topic
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			logrus.Errorf("error while reading message: %v", err)
		}

		fmt.Printf("message at offset %d: %s = %s\n\n", m.Offset, string(m.Key), string(m.Value))

		// Here, process the message as needed.
		// func process msg (unmarshall the msg value into a struct)
		// database func create log (create and save to db)
	}

	// in case of a finite read
	// return c.reader.Close()
}

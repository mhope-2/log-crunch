// Package kafka houses kafka components
package kafka

import (
	"context"
	"github.com/mhope-2/log-agent/shared"
	"github.com/segmentio/kafka-go"
)

const TOPIC = "topic"

// WriterInterface abstracts the capability to write Kafka messages.
type WriterInterface interface {
	WriteMessages(ctx context.Context, msgs ...kafka.Message) error
}

// Producer represents a Kafka message producer
type Producer struct {
	writer WriterInterface
}

func NewProducer() *Producer {
	EnvConfig := shared.NewEnvEnvConfig()

	w := &kafka.Writer{
		Addr:  kafka.TCP(EnvConfig.KafkaBrokers...),
		Topic: TOPIC,
	}
	return &Producer{
		writer: w,
	}
}

// NewProducerWithWriter provides a way to create a producer with a custom writer, useful for testing
func NewProducerWithWriter(writer WriterInterface) *Producer {
	return &Producer{
		writer: writer,
	}
}

// SendMessage sends a message to the Kafka topic.
func (p *Producer) SendMessage(ctx context.Context, key, value []byte) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
}

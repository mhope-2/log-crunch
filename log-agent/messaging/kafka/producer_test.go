package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// MockWriter provides a mock for the WriterInterface.
type MockWriter struct {
	mock.Mock
}

// WriteMessages is the mock method for writing messages to Kafka.
func (m *MockWriter) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	args := m.Called(ctx, msgs)
	return args.Error(0)
}

// TestSendMessage tests that Producer.SendMessage properly uses WriterInterface to send messages.
func TestSendMessage(t *testing.T) {
	mockWriter := new(MockWriter)

	key := []byte("key")
	value := []byte("value")
	msgs := []kafka.Message{{Key: key, Value: value}}

	// Set up the expected call to WriteMessages with any context and the prepared messages.
	// It's expected to succeed, hence Return(nil) for no error.
	mockWriter.On("WriteMessages", mock.Anything, msgs).Return(nil)

	producer := NewProducerWithWriter(mockWriter)

	// Call SendMessage, which under the hood, should use our mock's WriteMessages method.
	err := producer.SendMessage(context.Background(), key, value)

	// Assert no error was returned and the mock's expectations were met.
	assert.NoError(t, err)
	mockWriter.AssertExpectations(t)
}

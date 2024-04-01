package processing

import (
	"fmt"
	"strings"
)

type Message struct {
	Timestamp string
	LogLevel  string
	Message   string
}

// ProcessMessage unmarshalls a log message in to a struct
func ProcessMessage(value string) (*Message, error) {

	parts := strings.SplitN(value, " ", 3)
	if len(parts) < 3 {
		fmt.Println("Invalid log message format")
		return nil, fmt.Errorf("invalid log message format")
	}

	timestamp := parts[0]
	logLevel := parts[1]
	message := parts[2]

	msg := Message{
		Timestamp: timestamp,
		LogLevel:  logLevel,
		Message:   message,
	}

	return &msg, nil
}

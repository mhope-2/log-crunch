package processing

import (
	"fmt"
	"github.com/mhope-2/log-reader/database/types"
	"github.com/mhope-2/log-reader/database/utils"
	"strings"
)

// ProcessMessage unmarshalls a log message in to a struct
func ProcessMessage(logMessage string) (*types.Message, error) {

	parts := strings.SplitN(logMessage, " ", 3)
	if len(parts) < 3 {
		fmt.Println("Invalid log message format")
		return nil, fmt.Errorf("invalid log message format")
	}

	timestamp := parts[0]
	logLevel := parts[1]
	value := parts[2]

	msg := types.Message{
		Timestamp: utils.ParseTime(timestamp),
		LogLevel:  logLevel,
		Value:     value,
	}

	return &msg, nil
}

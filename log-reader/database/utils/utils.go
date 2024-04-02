package utils

import (
	"time"
)

func ParseTime(timeStr string) time.Time {
	// Parse the time string using the RFC3339 layout
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		// Return the zero value of time.Time and the error
		return time.Time{}
	}

	return parsedTime
}

package types

import "time"

type Message struct {
	ID        string
	LogLevel  string
	Value     string
	Timestamp time.Time
}

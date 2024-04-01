package types

import "time"

type Message struct {
	ID        string
	Level     string
	Value     string
	Timestamp time.Time
}

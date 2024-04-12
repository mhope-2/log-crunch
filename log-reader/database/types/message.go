package types

import "time"

type Message struct {
	ID        string    `cql:"id"`
	LogLevel  string    `cql:"log_level"`
	Value     string    `cql:"value"`
	Timestamp time.Time `cql:"timestamp"`
}

// Package shared maintains reference to Shared variables and types
package shared

import (
	"os"
)

type EnvConfig struct {
	KafkaBrokers []string
	Topic        string
	GroupID      string
}

func NewEnvEnvConfig() *EnvConfig {
	return &EnvConfig{
		KafkaBrokers: []string{
			os.Getenv("KAFKA_BROKER1"),
			os.Getenv("KAFKA_BROKER2"),
			os.Getenv("KAFKA_BROKER3"),
		},
		Topic:   os.Getenv("TOPIC"),
		GroupID: os.Getenv("GROUP_ID"),
	}
}

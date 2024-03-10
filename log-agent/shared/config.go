// Package shared maintains reference to Shared variables and types
package shared

import (
	"os"
)

type EnvConfig struct {
	KafkaBrokers []string
}

func NewEnvEnvConfig() *EnvConfig {
	return &EnvConfig{
		KafkaBrokers: []string{
			os.Getenv("KAFKA_BROKER1"),
			os.Getenv("KAFKA_BROKER2"),
			os.Getenv("KAFKA_BROKER3")},
	}
}

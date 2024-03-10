package shared

import (
	"os"
	"reflect"
	"testing"
)

func TestNewEnvEnvConfig(t *testing.T) {
	// Define the environment variables for testing
	testBrokers := []string{"kafka1:9092", "kafka1:9093", "kafka1:9094"}
	os.Setenv("KAFKA_BROKER1", testBrokers[0])
	os.Setenv("KAFKA_BROKER2", testBrokers[1])
	os.Setenv("KAFKA_BROKER3", testBrokers[2])

	defer func() {
		// Cleanup environment variables after test
		os.Unsetenv("KAFKA_BROKER1")
		os.Unsetenv("KAFKA_BROKER2")
		os.Unsetenv("KAFKA_BROKER3")
	}()

	// Call the function under test
	config := NewEnvEnvConfig()

	// Check if the result matches the expected outcome
	expected := &EnvConfig{KafkaBrokers: testBrokers}
	if !reflect.DeepEqual(config, expected) {
		t.Errorf("NewEnvEnvConfig() = %v, want %v", config, expected)
	}
}

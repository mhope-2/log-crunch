package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mhope-2/log-agent/messaging/kafka"
	"github.com/sirupsen/logrus"
)

func main() {
	// wait for brokers to be fully online
	time.Sleep(10 * time.Second)

	logger := logrus.New()

	producer := kafka.NewProducer()

	logLevels := []string{"INFO", "DEBUG", "ERROR"}
	logMessages := []string{
		"User login successful",
		"User attempted invalid login",
		"Database connection timeout",
		"New user registration",
		"Payment processed successfully",
		"Payment failed due to insufficient funds",
		"Configuration file updated",
		"Failed to connect to external API",
		"Cache cleared successfully",
		"Unexpected error processing user request",
	}

	for {
		key := "log"
		message := logMessages[rand.Intn(len(logMessages))]
		level := logLevels[rand.Intn(len(logLevels))]
		timestamp := time.Now().Format(time.RFC3339)

		value := fmt.Sprintf("%s %s %s", timestamp, level, message)

		fmt.Println("Sending log message")

		err := producer.SendMessage(context.Background(), []byte(key), []byte(value))
		if err != nil {
			logger.Fatalf("Error sending message to Kafka: %v", err)
		}

		fmt.Println("Log message sent")

		time.Sleep(1 * time.Second)
	}
}

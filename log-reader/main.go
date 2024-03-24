package main

import (
	"context"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mhope-2/log-reader/messaging/kafka"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	// wait for brokers to be fully online
	time.Sleep(15 * time.Second)

	logger := logrus.New()

	ctx := context.Background()
	consumer := kafka.NewConsumer()

	if err := consumer.Consume(ctx); err != nil {
		logger.Fatalf("failed to start consumer: %v", err)
	}
}

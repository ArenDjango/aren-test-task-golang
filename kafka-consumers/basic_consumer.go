package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	// Set up the Kafka reader with the appropriate group ID
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "my-group",     // specify the consumer group ID
		Topic:    "ws-6-example", // specify the topic
		MinBytes: 10e3,           // 10KB
		MaxBytes: 10e6,           // 10MB
	})

	defer r.Close()

	for {
		// Read messages
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	// Note: No need to manually manage offsets, as Kafka takes care of it with consumer groups
}

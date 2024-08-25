package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

// with choosing partition
//func main() {
//	// to produce messages
//	topic := "my-topic"
//	partition := 1
//
//	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9091", topic, partition)
//	if err != nil {
//		log.Fatal("failed to dial leader:", err)
//	}
//
//	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
//	_, err = conn.WriteMessages(
//		kafka.Message{Value: []byte("one 1!")},
//		kafka.Message{Value: []byte("two 1!")},
//		kafka.Message{Value: []byte("three 1!")},
//	)
//	if err != nil {
//		log.Fatal("failed to write messages:", err)
//	}
//
//	if err := conn.Close(); err != nil {
//		log.Fatal("failed to close writer:", err)
//	}
//}

// with rund-roin
func main() {
	w := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092", "localhost:9093"),
		Topic:        "ws-6-example",
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: []byte("Hello World!"),
		},
		kafka.Message{
			Key:   []byte("Key-B"),
			Value: []byte("One!"),
		},
		kafka.Message{
			Key:   []byte("Key-C"),
			Value: []byte("Two!"),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

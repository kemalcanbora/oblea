package redpanda

import (
	"github.com/segmentio/kafka-go"
	"time"
)

func RedProducer() {
	client.SetWriteDeadline(time.Now().Add(10 * time.Second))
	client.WriteMessages(
		kafka.Message{Value: []byte("goodbye world")},
	)
}

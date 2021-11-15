package redpanda

import (
	"context"
	"github.com/segmentio/kafka-go"
)

func RedClient() *kafka.Conn {
	ctx := context.Background()
	conn, err := kafka.DialLeader(ctx, "tcp", "localhost:9092", "test", 0)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	return conn
}

var client = RedClient()

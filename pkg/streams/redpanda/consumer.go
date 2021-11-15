package redpanda

import (
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func RedConsumer() *kafka.Conn {
	client.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := RedClient().ReadBatch(10e3, 1e6) //fetch 10KB min, 1MB max
	b := make([]byte, 10e3)                   //10kb max per message
	for {
		n, err := batch.Read(b)
		if err != nil {
			break
		}
		fmt.Println(string(b[:n]))
	}
	//close
	if err := RedClient().Close(); err != nil {
		panic(err)
	}
	return nil
}

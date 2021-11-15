package redpanda

import (
	"fmt"
	panda "github.com/segmentio/kafka-go"
	"oblea/config"
)

func PandaReader(topic string) *panda.Reader {
	address := fmt.Sprintf("%s:%d", config.Configure().RedPanda.Host, config.Configure().RedPanda.Port)
	r := panda.NewReader(panda.ReaderConfig{
		Brokers:   []string{address},
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	return r
}

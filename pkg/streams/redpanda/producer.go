package redpanda

import (
	"fmt"
	panda "github.com/segmentio/kafka-go"
	"oblea/config"
)

func PandaWriter(topic string) panda.Writer {
	address := fmt.Sprintf("%s:%d", config.Configure().RedPanda.Host, config.Configure().RedPanda.Port)
	return panda.Writer{
		Addr:     panda.TCP(address),
		Topic:    topic,
		Balancer: &panda.LeastBytes{},
	}
}

// Message funciton
func Message(msg string) panda.Message {
	pandaMsg := panda.Message{
		Value: []byte(msg),
	}
	return pandaMsg
}

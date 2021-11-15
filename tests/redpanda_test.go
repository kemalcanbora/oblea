package main

import (
	"context"
	"oblea/pkg/streams/redpanda"
	"testing"
)

var ctx = context.Background()

func TestWriteMessage(t *testing.T) {
	writer := redpanda.PandaWriter("topic-A")
	wErr := writer.WriteMessages(ctx, redpanda.Message("hello world"))
	if wErr != nil {
		t.Errorf("We have a problem %s", wErr)
	}
}

func TestReadMessage(t *testing.T) {
	reader := redpanda.PandaReader("topic-A")
	message, rErr := reader.ReadMessage(ctx)
	if string(message.Value) != "hello world" {
		t.Errorf("We have a problem %s", rErr)
	}
}

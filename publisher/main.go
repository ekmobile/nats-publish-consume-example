package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	url := "localhost:4222"
	nc, _ := nats.Connect(url)
	js, _ := jetstream.New(nc)

	js.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})

	for i := 0; i < 10; i++ {
		js.Publish(ctx, "ORDERS.new", []byte("hello message "+strconv.Itoa(i)))
		fmt.Printf("Published hello message %d\n", i)
	}
}

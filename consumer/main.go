package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var messageCounter = 0

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	url := "localhost:4222"
	nc, _ := nats.Connect(url)
	js, _ := jetstream.New(nc)

	c, error := js.CreateOrUpdateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
		Durable:   "CONS",
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if error != nil {
		panic(error)
	}

	cons, err := c.Consume(handleConsume, jetstream.ConsumeErrHandler(handleError))
	if err != nil {
		panic(err)
	} else {
		log.Println("cons", cons)
	}
	defer cons.Stop()

	// block until all 100 published messages have been processed
	for messageCounter < 100 {
		time.Sleep(10 * time.Millisecond)
	}
}

func handleError(consumeCtx jetstream.ConsumeContext, err error) {
	println("handleError")
	fmt.Printf("Received a JetStream error via callback: %s\n", string(err.Error()))
}

func handleConsume(msg jetstream.Msg) {
	println("handleConsume")
	messageCounter++
	fmt.Printf("Received a JetStream message via callback: %s\n", string(msg.Data()))
	msg.Ack()
}

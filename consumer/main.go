package main

import (
	"context"
	"fmt"
	"log"
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

	ordersStream, err := js.Stream(ctx, "ORDERS")
	if err != nil {
		log.Fatalf(err.Error())
	}

	ordersConsumer, err := ordersStream.Consumer(ctx, "cons")
	if err != nil {
		log.Fatalf("Could not setup consumer: %s", err.Error())
	}

	cc, err := ordersConsumer.Consume(handleConsume, jetstream.ConsumeErrHandler(handleError))
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer cc.Stop()

	for {

	}
}

func handleError(consumeCtx jetstream.ConsumeContext, err error) {
	println("handleError")
	fmt.Printf("Received a JetStream error via callback: %s\n", string(err.Error()))
}

func handleConsume(msg jetstream.Msg) {
	println("handleConsume")
	fmt.Printf("Received a JetStream message via callback: %s\n", string(msg.Data()))
	msg.Ack()
}

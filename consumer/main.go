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

	c, _ := js.CreateOrUpdateConsumer(ctx, "ORDERS", jetstream.ConsumerConfig{
		Name:           "CONS",
		Durable:        "CONS",
		AckPolicy:      jetstream.AckExplicitPolicy,
		FilterSubjects: []string{"ORDERS.*"},
	})

	cons, err := c.Consume(handleConsume, jetstream.ConsumeErrHandler(handleError))
	if err != nil {
		panic(err)
	} else {
		log.Println(cons)
	}
	defer cons.Stop()
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

package main

import (
	"encoding/json"
	"fmt"

	"github.com/gstanleysilva/fc-go-events-rmq/internal/domain"
	"github.com/gstanleysilva/fc-go-events-rmq/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	//Consume messages from the orders queue
	go rabbitmq.Consume(ch, msgs, "orders")

	for msg := range msgs {
		//Convert to the Order Structure
		var order domain.Order
		json.Unmarshal(msg.Body, &order)

		//Print our result
		fmt.Println(order)

		//Notify RMQ that we received the message
		msg.Ack(false)
	}
}

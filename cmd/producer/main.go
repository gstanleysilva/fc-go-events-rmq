package main

import (
	rmqImp "github.com/gstanleysilva/fc-go-events-rmq/infra/rabbitmq"
	"github.com/gstanleysilva/fc-go-events-rmq/internal/domain"
	"github.com/gstanleysilva/fc-go-events-rmq/pkg/events"
	"github.com/gstanleysilva/fc-go-events-rmq/pkg/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

var ch *amqp091.Channel

func main() {
	var err error

	ch, err = rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	//Create the object of a handler implementation
	RMQHandler := rmqImp.NewHandler(ch, "amq.direct")
	//Create the object of a dispatcher implementation
	dispatcher := events.NewEventDispatcher()
	//Register the handler for the OrderCreated event
	dispatcher.Register("OrderCreated", RMQHandler)

	CreateOrder(dispatcher)
}

func CreateOrder(dispatcher events.Dispatcher) {
	//Execute methods to create an order
	//...
	order := domain.Order{
		ID:         "1",
		CustomerID: "1",
		Items: []domain.Item{
			{
				ID:    "1",
				Title: "Item 1",
				Price: 10.99,
			},
		},
	}

	dispatcher.Dispatch(events.NewEvent("OrderCreated", order))
}

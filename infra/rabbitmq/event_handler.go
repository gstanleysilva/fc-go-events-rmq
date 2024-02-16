package rabbitmq

import (
	"encoding/json"
	"sync"

	"github.com/gstanleysilva/fc-go-events-rmq/pkg/events"
	"github.com/gstanleysilva/fc-go-events-rmq/pkg/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

type Handler struct {
	ch     *amqp091.Channel
	exName string
}

func NewHandler(ch *amqp091.Channel, exName string) *Handler {
	return &Handler{
		ch:     ch,
		exName: exName,
	}
}

func (m *Handler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	//Convert to JSON
	payload, err := json.Marshal(event.GetPayload())
	if err != nil {
		panic(err)
	}
	//Publish the json to the exchange
	rabbitmq.Publish(m.ch, payload, m.exName)
	wg.Done()
}

package client

import (
	"context"
	"errors"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/wenealves10/tracking-app-golang/domain"
)

const (
	QueueName = "package_status"
)

type rabbitmqCliet struct {
	conn         *amqp.Connection
	ch           *amqp.Channel
	packageStaus <-chan amqp.Delivery
}

func NewRabbitClient(connString string) (*rabbitmqCliet, error) {
	c := &rabbitmqCliet{}

	var err error

	c.conn, err = amqp.Dial(connString)

	if err != nil {
		return nil, err
	}

	err = c.configureQueue()

	if err != nil {
		return nil, err
	}

	return c, err

}

func (rb *rabbitmqCliet) Publish(p domain.Package) error {
	jsonStr := fmt.Sprintf(`{ "from": %q, "to": %q, "vehicleID": %q }`, p.From, p.To, p.VehicleID)

	err := rb.ch.Publish("", QueueName, true, false, amqp.Publishing{
		ContentType: "application/json",
		MessageId:   p.VehicleID,
		Body:        []byte(jsonStr),
	})

	if err != nil {
		return err
	}

	return err
}

func (rb *rabbitmqCliet) ConsumeByVehicleID(ctx context.Context, vehicleID string) ([]byte, error) {
	for msg := range rb.packageStaus {
		if msg.MessageId == vehicleID {
			_ = msg.Ack(false)
			return msg.Body, nil
		}
	}

	return nil, errors.New("err when getting package status on channel")
}

func (rb *rabbitmqCliet) Close() {
	err := rb.ch.Close()

	if err != nil {
		panic(err)
	}

	err = rb.conn.Close()

	if err != nil {
		panic(err)
	}

}

func (rb *rabbitmqCliet) configureQueue() error {
	_, err := rb.ch.QueueDeclare(
		QueueName,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	rb.packageStaus, err = rb.ch.Consume(
		QueueName,
		"",
		false,
		false,
		false,
		true,
		nil,
	)

	return err

}

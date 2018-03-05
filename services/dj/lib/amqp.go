package lib

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/streadway/amqp"

	"qcast"
)

func exchangeName(party string) string {
	return fmt.Sprintf("party_%s", party)
}

func QueueUpdates(party string, conn *amqp.Connection) (chan *qcast.QueueUpdate, error) {
	// open up an amqp channel
	ch, err := conn.Channel()
	if err != nil {
		glog.Errorf("Error opening channel %s", err)
		return nil, err
	}

	updates := make(chan *qcast.QueueUpdate)

	err = ch.ExchangeDeclare(
		exchangeName(party), // name
		"fanout",            // type
		true,                // durable
		true,                // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		glog.Errorf("Error declaring exchange %s", err)
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		true,  // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		glog.Errorf("Error declaring queue %s", err)
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,              // queue name
		"",                  // routing key
		exchangeName(party), // exchange
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		glog.Errorf("Error binding queue %s", err)
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		glog.Errorf("Error consuming from queue %s", err)
		return nil, err
	}

	go func() {
		defer ch.Close()
		for msg := range msgs {
			update := qcast.QueueUpdate{}
			err = proto.Unmarshal(msg.Body, &update)
			if err != nil {
				glog.Errorf("Error unmarshalling queue update: %s", err)
				continue
			}

			updates <- &update
		}
	}()

	return updates, nil
}

func SendToRabbitmq(party string, update *qcast.QueueUpdate, conn *amqp.Connection) error {
	ch, err := conn.Channel()
	if err != nil {
		glog.Errorf("Error opening channel: %s", err)
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName(party), // name
		"fanout",            // type
		true,                // durable
		true,                // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		glog.Errorf("Error declaring exchange: %s", err)
		return err
	}

	body, err := proto.Marshal(update)
	if err != nil {
		glog.Errorf("Error marshaling update: %s", err)
		return err
	}

	err = ch.Publish(
		exchangeName(party), // exchange
		"",                  // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "application/octet-stream",
			Body:        []byte(body),
		},
	)
	if err != nil {
		glog.Errorf("Error publishing update: %s", err)
		return err
	}

	return nil
}

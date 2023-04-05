package event

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	amqp "github.com/rabbitmq/amqp091-go"
// )

// type Consumers struct {
// 	conn      *amqp.Connection
// 	queueName string
// }

// func NewConsumer(conn *amqp.Connection) (Consumers, error) {
// 	consumers := Consumers{
// 		conn: conn,
// 	}

// 	err := consumers.setup()
// 	if err != nil {
// 		return Consumers{}, err
// 	}

// 	return consumers, nil
// }

// type Payload struct {
// 	Name string `json:"name"`
// 	Data string `json:"data"`
// }

// func (consumers *Consumers) setup() error {
// 	channel, err := consumers.conn.Channel()
// 	if err != nil {
// 		return err
// 	}

// 	return declareExcange(channel)
// }

// func (consumers *Consumers) Listen(topic []string) error {
// 	ch, err := consumers.conn.Channel()
// 	if err != nil {
// 		return err
// 	}
// 	defer ch.Close()

// 	queue, err := declareRandomQueue(ch)
// 	if err != nil {
// 		return err
// 	}

// 	for _, s := range topic {
// 		ch.QueueBind(
// 			queue.Name,
// 			s,
// 			"logs_topic",
// 			false,
// 			nil,
// 		)

// 		if err != nil {
// 			return err
// 		}
// 	}

// 	message, err := ch.Consume(
// 		queue.Name,
// 		"",
// 		true,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)

// 	forever := make(chan bool)
// 	go func() {
// 		for d := range message {
// 			var payload Payload
// 			_ = json.Unmarshal(d.Body, &payload)

// 			go handlePayload(payload)
// 		}
// 	}()

// 	fmt.Printf("Waiting for message [Exchange, Queue] [logs_topic, %s]\n", queue.Name)
// 	<-forever

// 	return nil
// }

package event

import (
	"github.com/streadway/amqp"
)

func ExchangeDeclare(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"kredit-exchange", // Nama exchange
		"topic",           // Tipe exchange
		true,              // Durable
		false,             // Auto-deleted
		false,             // Internal
		false,             // No-wait
		nil,               // Arguments
	)
}

func Publish(ch *amqp.Channel, message []byte) error {
	// Mempublish message ke exchange dengan routing key "test-topic"
	err := ch.Publish(
		"kredit-exchange", // Nama exchange
		"transaksi-topic", // Routing key
		false,             // Mandatory
		false,             // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)

	return err
}

func QueueDeclare(ch *amqp.Channel) (amqp.Queue, error) {
	// Mendeklarasikan queue
	q, err := ch.QueueDeclare(
		"",    // Nama queue
		false, // Durable
		false, // Delete when unused
		true,  // Exclusive
		false, // No-wait
		nil,   // Arguments
	)

	return q, err
}

func QueueBind(ch *amqp.Channel, queue amqp.Queue) error {
	err := ch.QueueBind(
		queue.Name,        // Nama queue
		"transaksi-topic", // Routing key
		"kredit-exchange", // Nama exchange
		false,
		nil,
	)

	return err
}

func Consume(ch *amqp.Channel, queue amqp.Queue) (<-chan amqp.Delivery, error) {
	msgs, err := ch.Consume(
		queue.Name, // Nama queue
		"",         // Consumer
		true,       // Auto-acknowledge
		false,      // Exclusive
		false,      // No-local
		false,      // No-wait
		nil,        // Arguments
	)

	return msgs, err
}

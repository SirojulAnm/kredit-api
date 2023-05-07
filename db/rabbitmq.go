package db

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func ConnectRMQ() (*amqp.Connection, error) {
	connectString := os.Getenv("RMQ")
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// dont continue until rabbit is ready
	for {
		conn, err := amqp.Dial(connectString)
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ!")
			connection = conn
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("Backing off...")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}

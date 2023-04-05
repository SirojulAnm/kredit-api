package event

// func InitRabbitMQ() (*amqp.Channel, string, error) {
// 	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	// defer conn.Close()
// 	fmt.Println(conn)
// 	ch, err := conn.Channel()
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	defer ch.Close()
// 	q, err := ch.QueueDeclare(
// 		"queue_name",
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	return ch, q.Name, nil
// }

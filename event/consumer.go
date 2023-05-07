package event

import (
	"encoding/json"
	"kredit-api/db"
	"kredit-api/transaksi"
	"log"

	"gorm.io/gorm"
)

func ListenTransaksi(gormDB *gorm.DB) {
	var input transaksi.TransaksiInput

	transaksiRepository := transaksi.NewRepository(gormDB)
	transaksiService := transaksi.NewService(transaksiRepository)

	rabbitConn, err := db.ConnectRMQ()
	if err != nil {
		// log.Fatal(err)
		log.Fatalf("Error connect RabbitMQ, %v", err)
	}
	defer rabbitConn.Close()

	ch, err := rabbitConn.Channel()
	if err != nil {
		log.Fatalf("Error Create Channel Transaksi, %v", err)
	}

	log.Println("Listening for and consuming RabbitMQ messages...")

	err = ExchangeDeclare(ch)
	if err != nil {
		// return err
		log.Fatalf("Failed to ExchangeDeclare: %s", err)
	}

	queue, err := QueueDeclare(ch)
	if err != nil {
		// return err
		log.Fatalf("Failed to QueueDeclare: %s", err)
	}

	err = QueueBind(ch, queue)

	ready := make(chan bool)

	// go func() {
	msgs, err := Consume(ch, queue)

	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	for {
		select {
		case msg := <-msgs:
			log.Printf("Received a message: %s", msg.Body)
			err := json.Unmarshal(msg.Body, &input)
			if err != nil {
				log.Printf("Failed to deserialize message body: %v", err.Error())
				continue
			}

			newTransaksi, err := transaksiService.AddTransaksi(input)
			if err != nil {
				log.Printf("failed saat insert db: %v", err.Error())
				continue
			}

			log.Printf("Success Add transaksi: %v", newTransaksi)
		case <-ready:
			log.Println("Consumer is ready")
		}
	}
	// }()

	// ready <- true

	// select {}
}

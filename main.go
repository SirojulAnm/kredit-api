package main

import (
	"context"
	"kredit-api/db"
	"kredit-api/event"
	"kredit-api/router"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	mongoClient, err := db.MongoConnect()
	if err != nil {
		panic(err)
	}
	defer mongoClient.Disconnect(context.Background())

	gormDB, err := db.Open()
	if err != nil {
		log.Fatal(err)
	}

	go event.ListenTransaksi(gormDB)

	err = router.Router(gormDB, mongoClient)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for a signal to gracefully shutdown the application
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	log.Println("Shutting down...")
}

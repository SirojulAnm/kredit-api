package db

import (
	"context"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect() (*mongo.Client, error) {
	connectString := os.Getenv("MONGO")
	clientOptions := options.Client().ApplyURI(connectString)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Cek koneksi ke MongoDB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	log.Println("Berhasil terkoneksi ke MongoDB!")
	return client, nil
}

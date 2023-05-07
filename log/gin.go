package log

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type RequestLog struct {
	Method     string
	Path       string
	StatusCode int
	Latency    time.Duration
	CreatedAt  time.Time
}

func LogGin(c *gin.Context, mongoClient *mongo.Client) {
	// Simpan waktu awal permintaan
	startTime := time.Now()

	// Handle permintaan
	c.Next()

	// Buat struct log
	requestLog := RequestLog{
		Method:     c.Request.Method,
		Path:       c.Request.URL.Path,
		StatusCode: c.Writer.Status(),
		Latency:    time.Duration(time.Since(startTime).Milliseconds()),
		CreatedAt:  time.Now(),
	}

	// Simpan log ke MongoDB
	data, err := mongoClient.Database("kredit").Collection("gin_logs").InsertOne(context.Background(), requestLog)
	if err != nil {
		log.Println("Failed to insert log to MongoDB: ", err)
	}
	log.Println("Succes insert data mongo", data)
}

package router

import (
	"kredit-api/log"
	"kredit-api/router/group"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

func Router(DB *gorm.DB, mongoClient *mongo.Client) error {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Static("/images", "./images")
	corsConfig(router)

	router.Use(func(ctx *gin.Context) {
		log.LogGin(ctx, mongoClient)
	})

	v1 := router.Group("v1")
	group.AuthRouter(DB, v1)
	group.TransaksiRouter(DB, v1)

	err := router.Run("0.0.0.0:8000")
	if err != nil {
		return err
	}

	return nil
}

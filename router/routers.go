package router

import (
	"kredit-api/router/group"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Router(DB *gorm.DB) error {
	router := gin.Default()
	router.Static("/images", "./images")
	corsConfig(router)

	v1 := router.Group("v1")
	group.AuthRouter(DB, v1)
	group.TransaksiRouter(DB, v1)

	err := router.Run(":8000")
	if err != nil {
		return err
	}

	return nil
}

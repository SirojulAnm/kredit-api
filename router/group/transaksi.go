package group

import (
	"kredit-api/auth"
	"kredit-api/handler"
	"kredit-api/transaksi"
	"kredit-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TransaksiRouter(db *gorm.DB, main *gin.RouterGroup) {
	userRepository := user.NewRepository(db)
	transaksiRepository := transaksi.NewRepository(db)

	authService := auth.NewService()
	userService := user.NewService(userRepository)
	transaksiService := transaksi.NewService(transaksiRepository)

	transaksiHandler := handler.NewTransaksiHandler(transaksiService, userService)
	userHandler := handler.NewUserHandler(userService, authService)

	transaksiGroup := main.Group("/transaksi")
	{
		transaksiGroup.POST("/add", authMiddleware(userService, authService), transaksiHandler.AddTransaksi)
		transaksiGroup.GET("/history", authMiddleware(userService, authService), userHandler.HistoryTransaksi)
	}
}

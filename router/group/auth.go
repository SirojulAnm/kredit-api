package group

import (
	"fmt"
	"kredit-api/auth"
	"kredit-api/handler"
	"kredit-api/helper"
	"kredit-api/user"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRouter(db *gorm.DB, main *gin.RouterGroup) {
	userRepository := user.NewRepository(db)

	authService := auth.NewService()
	userService := user.NewService(userRepository)

	userHandler := handler.NewUserHandler(userService, authService)

	main.POST("/login", userHandler.Login)
	main.POST("/register", userHandler.Register)
	main.POST("/upload", authMiddleware(userService, authService), userHandler.UploadPhoto)
	main.GET("/profile", authMiddleware(userService, authService), userHandler.Profile)
}

func authMiddleware(userService user.Service, authable auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		fmt.Println(tokenString)
		if tokenString == "" {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		customClaim, err := authable.ValidateToken(tokenString)
		if customClaim == nil && err != nil {
			response := helper.APIResponse("SessionExpired", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID, err := strconv.ParseInt(customClaim.Subject, 10, 64)

		user, err := userService.GetUserByID(int(userID))
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

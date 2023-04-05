package handler

import (
	"fmt"
	"kredit-api/auth"
	"kredit-api/helper"
	"kredit-api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) Register(ctx *gin.Context) {
	var inputRegister user.InputRegister

	err := ctx.ShouldBindJSON(&inputRegister)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register failed saat input json", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.Register(inputRegister)

	if err != nil {
		response := helper.APIResponse("Register failed saat insert", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed generate token", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "")

	response := helper.APIResponse("Successfully register", http.StatusOK, "success", gin.H{"user": formatter, "token": token})

	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(ctx *gin.Context) {
	var input user.LoginAdminRequest

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login failed saat input json", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login failed saat cek email atau password", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedinUser.ID)
	if err != nil {
		response := helper.APIResponse("Login failed saat generate token", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatLogin(loggedinUser)

	response := helper.APIResponse("Success Log In", http.StatusOK, "success", gin.H{"user": formatter, "token": token})

	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) Profile(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)

	baseURL := ctx.Request.Host
	protocol := "http"
	if ctx.Request.TLS != nil {
		protocol = "https"
	}
	currentURL := fmt.Sprintf("%s://%s", protocol, baseURL)

	formatter := user.FormatUser(currentUser, currentURL)

	response := helper.APIResponse("Successfully fetch user data", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) HistoryTransaksi(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)

	userTransaksi, err := h.userService.GetHistoryTransaksiByUserID(currentUser.ID)
	if err != nil {
		response := helper.APIResponse("Error get data user", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	baseURL := ctx.Request.Host
	protocol := "http"
	if ctx.Request.TLS != nil {
		protocol = "https"
	}
	currentURL := fmt.Sprintf("%s://%s", protocol, baseURL)

	formatter := user.FormatHistoryTransaksi(userTransaksi, currentURL)

	response := helper.APIResponse("Success get data user transaksi", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadPhoto(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)

	//file KTP
	fileKtp, err := ctx.FormFile("file_ktp")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image ktp saat input file", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	pathKtp := fmt.Sprintf("images/%d/%s", currentUser.ID, fileKtp.Filename)
	err = ctx.SaveUploadedFile(fileKtp, pathKtp)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload image ktp saat save file", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	//file Selfie
	fileSelfie, err := ctx.FormFile("file_selfie")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image selfie saat input file", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	pathSelfie := fmt.Sprintf("images/%d/%s", currentUser.ID, fileSelfie.Filename)
	err = ctx.SaveUploadedFile(fileSelfie, pathSelfie)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("failed to upload image selfie saat save file", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.CreatePhoto(currentUser.ID, pathKtp, pathSelfie)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image name ke db", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("ktp, selfie image successfuly uploaded", http.StatusOK, "success", nil)

	ctx.JSON(http.StatusOK, response)
}

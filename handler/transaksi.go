package handler

import (
	"encoding/json"
	"kredit-api/db"
	"kredit-api/event"
	"kredit-api/helper"
	"kredit-api/transaksi"
	"kredit-api/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transaksiHandler struct {
	transaksiService transaksi.Service
	userService      user.Service
}

func NewTransaksiHandler(transaksiService transaksi.Service, userService user.Service) *transaksiHandler {
	return &transaksiHandler{transaksiService, userService}
}

func (h *transaksiHandler) AddTransaksi(ctx *gin.Context) {
	var input transaksi.TransaksiInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorsMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Add transaksi failed saat input json", http.StatusUnprocessableEntity, "error", errorsMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	conn, err := db.ConnectRMQ()
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	err = event.ExchangeDeclare(ch)
	if err != nil {
		log.Fatalf("Failed to declare an exchange: %v", err)
	}

	currentUser := ctx.MustGet("currentUser").(user.User)
	input.UserID = currentUser.ID

	messageBody, err := json.Marshal(input)
	err = event.Publish(ch, messageBody)
	if err != nil {
		response := helper.APIResponse("Add transaksi failed saat send RabbitMQ", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success Add transaksi", http.StatusOK, "success", "berhasil kirim rabbitmq")

	ctx.JSON(http.StatusOK, response)
}

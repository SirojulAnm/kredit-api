package handler

import (
	"encoding/json"
	"kredit-api/helper"
	"kredit-api/transaksi"
	"kredit-api/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
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

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a RabbitMQ channel: %v", err)
	}
	defer channel.Close()

	queue, err := channel.QueueDeclare(
		"myqueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a RabbitMQ queue: %v", err)
	}

	messageBody, err := json.Marshal(input)
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        messageBody,
	}

	err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		message,
	)
	if err != nil {
		log.Fatalf("Failed to publish message to RabbitMQ: %v", err)
	}

	msgs, err := channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to start consuming messages from RabbitMQ: %v", err)
	}

	var response helper.Response
	for msg := range msgs {
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			log.Printf("Failed to deserialize message body: %v", err)
			continue
		}

		currentUser := ctx.MustGet("currentUser").(user.User)
		userID := currentUser.ID

		input.UserID = userID
		newTransaksi, err := h.transaksiService.AddTransaksi(input)
		if err != nil {
			response := helper.APIResponse("Add transaksi failed saat insert db", http.StatusBadRequest, "error", err.Error())
			ctx.JSON(http.StatusBadRequest, response)
			return
		}

		formatter := transaksi.FormatTransaksi(newTransaksi)

		response = helper.APIResponse("Success Add transaksi", http.StatusOK, "success", formatter)

		ctx.JSON(http.StatusOK, response)
		return
	}

}

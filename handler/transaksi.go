package handler

import (
	"kredit-api/helper"
	"kredit-api/transaksi"
	"kredit-api/user"
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

	response := helper.APIResponse("Success Add transaksi", http.StatusOK, "success", formatter)

	ctx.JSON(http.StatusOK, response)
}

package helper

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	JsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return JsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func GetCurrentUrl(ctx *gin.Context) string {
	baseURL := ctx.Request.Host
	protocol := "http"

	if ctx.Request.TLS != nil {
		protocol = "https"
	}
	currentURL := fmt.Sprintf("%s://%s", protocol, baseURL)

	return currentURL
}

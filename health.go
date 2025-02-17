package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Response
	Code string `json:"code"`
}

// @BasePath /
// @Version 1.0
// @Tags Internal
// @Summary Check Health
// @Description Its used to know if the server is working correctly
// @Router /check-health [get]
// @Success 200 {object} Response "OK"
// @Failure 404 {object} ErrorResponse "Not found"
func CheckHealth(ctx *gin.Context) {
	response := Response{"OK"}
	ctx.JSON(http.StatusOK, response)
}

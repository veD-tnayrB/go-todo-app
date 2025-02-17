package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *TodoHandler) List(ctx *gin.Context) {
	todos, err := h.TodoService.List()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": todos})
}

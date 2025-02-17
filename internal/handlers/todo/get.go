package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *TodoHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "ID_REQUIRED"})
	}

	todo, err := h.TodoService.Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": todo})
}

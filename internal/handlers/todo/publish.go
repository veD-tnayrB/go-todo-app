package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/veD-tnayrB/todo-app/common/models"
)

func (h *TodoHandler) Create(ctx *gin.Context) {
	body := models.Todo{}
	if err := ctx.ShouldBindBodyWithJSON(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "malformed body"})
		return
	}

	if body.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "and the tilte?"})
	}

	err := h.TodoService.Create(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok ma fren, you did it, perfext"})
}

func (h *TodoHandler) Update(ctx *gin.Context) {
	body := models.Todo{}
	if err := ctx.ShouldBindBodyWithJSON(body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "malformed body"})
		return
	}

	if body.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "and the tilte?"})
	}

	err := h.TodoService.Update(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "bad"})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "ok ma fren, you did it, perfext but this time updating wich its harder"})
}

func (h *TodoHandler) Remove(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "and the id yoou fucking monke?"})
	}

	err := h.TodoService.Remove(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "bad, your faul."})
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "al right"})
}

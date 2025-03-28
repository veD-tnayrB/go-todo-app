package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/veD-tnayrB/todo-app/common/models"
	"github.com/veD-tnayrB/todo-app/common/responses"
)

// @Summary Todo List
// @Description Lists all the todos existing in the DB without filters
// @BasePath /todos
// @Router /todos [get]
// @Version 1.0
// @Tags todos
// @Success 200 {object} responses.Success[[]models.Todo]
// @Failure 500 {object} responses.Error
func (h *TodoHandler) List(ctx *gin.Context) {
	todos, err := h.TodoService.List()
	h.Logger.Info("Handler: Handling LIST request")
	if err != nil {
		h.Logger.Error("Handler: Error Handling LIST request", "error", err)
		ctx.JSON(http.StatusInternalServerError, responses.Error{Status: false, Code: 500, Message: "SOMETHING_WENT_WRONG"})
		return
	}

	response := responses.Success[*[]*models.Todo]{
		Status: true,
		Data:   todos,
	}

	ctx.JSON(http.StatusOK, response)
}

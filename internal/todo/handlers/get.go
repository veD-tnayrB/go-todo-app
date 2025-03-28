package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/veD-tnayrB/todo-app/common/models"
	"github.com/veD-tnayrB/todo-app/common/responses"
)

// @Summary Todo get
// @Description Allows you to get one item by its id
// @Router /todos/:id [get]
// @Param string query string true "string valid" maxlength(36)
// @Version 1.0
// @Tags todos
// @Success 200 {object} responses.Success[models.Todo]
// @Failure 400 {object} responses.Error "Id is required"
// @Failure 500 {object} responses.Error "Something went wrong"
func (h *TodoHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	h.Logger.Info("Handler: Handling GET request", "id", id)

	if id == "" {
		message := ErrIdIsRequired.Error()
		h.Logger.Warn("Handler: Missing id in request", "error", message)
		err := responses.Error{Status: false, Code: 400, Message: message}
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	todo, err := h.TodoService.Get(id)
	if err != nil {
		h.Logger.Error("Handler: Todo get service failed", "error", err.Error())
		err := responses.Error{Status: false, Code: 500, Message: err.Error()}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, responses.Success[*models.Todo]{Status: true, Data: todo})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/veD-tnayrB/todo-app/common/models"
	"github.com/veD-tnayrB/todo-app/common/responses"
	"github.com/veD-tnayrB/todo-app/internal/errors"
)

// @Summary Todo get
// @Description Allows you to get one item by its id
// @Router /todo/get/:id [get]
// @Param string query string true "string valid" maxlength(36)
// @Version 1.0
// @Tags todos
// @Success 200 {object} responses.Success[models.Todo]
// @Failure 400 {object} responses.Error "Id is required"
// @Failure 500 {object} responses.Error "Something went wrong"
func (h *TodoHandler) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		err := responses.Error{Status: false, Code: 400, Message: errors.IdIsRequired.Error()}
		ctx.JSON(http.StatusBadRequest, err)
	}

	todo, err := h.TodoService.Get(id)
	if err != nil {
		err := responses.Error{Status: false, Code: 500, Message: err.Error()}
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, responses.Success[*models.Todo]{Status: true, Data: todo})
}

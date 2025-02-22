package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/veD-tnayrB/todo-app/common/models"
	"github.com/veD-tnayrB/todo-app/common/responses"
	"github.com/veD-tnayrB/todo-app/internal/errors"
)

// @Summary Create todo
// @Description Allows you to create a todo by passing your todo structure
// @Router /todo/create [post]
// @Version 1.0
// @Tags todos
// @Param todo body models.Todo true "Todo details"
// @Success 200 {object} responses.Empty
// @Failure 400 {object} responses.Error "Malformed body"
// @Failure 400 {object} responses.Error "Title is required"
// @Failure 500 {object} responses.Error "Something went wrong"
func (h *TodoHandler) Create(ctx *gin.Context) {
	body := models.Todo{}
	if err := ctx.ShouldBindBodyWithJSON(body); err != nil {
		ctx.JSON(http.StatusBadRequest, errors.MalformedBody)
		return
	}

	if body.Title == "" {
		ctx.JSON(http.StatusBadRequest, errors.TitleIsRequired)
	}

	err := h.TodoService.Create(&body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.Success[models.Todo]{Status: true, Data: body})
	}

	ctx.JSON(http.StatusOK, responses.Empty{Status: true})
}

// @Summary Updating todo
// @Description Allows you to update a todo by passing its id and the todo structure
// @Router /todo/update [put]
// @Version 1.0
// @Tags todos
// @Param todo body models.Todo true "Todo details"
// @Success 200 {object} responses.Empty
// @Failure 400 {object} responses.Error "Malformed body"
// @Failure 400 {object} responses.Error "Title is required"
// @Failure 400 {object} responses.Error "Id is required"
// @Failure 500 {object} responses.Error "Something went wrong"
func (h *TodoHandler) Update(ctx *gin.Context) {
	body := models.Todo{}
	if err := ctx.ShouldBindBodyWithJSON(body); err != nil {
		ctx.JSON(http.StatusBadRequest, errors.MalformedBody)
		return
	}

	if body.Title == "" {
		ctx.JSON(http.StatusBadRequest, errors.TitleIsRequired)
	}

	err := h.TodoService.Update(&body)
	if err != nil {
		if err == errors.RecordNotExists {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusInternalServerError, errors.ErrorSaving)

	}

	ctx.JSON(http.StatusOK, responses.Empty{Status: true})
}

// @Summary Removing todo
// @Description Allows you to remove a todo by passing its id
// @Router /todo/delete [delete]
// @Version 1.0
// @Tags todos
// @Param id identifier string true "Todo id"
// @Success 200 {object} responses.Empty
// @Failure 400 {object} responses.Error "Id is required"
// @Failure 500 {object} responses.Error "Record not exists"
// @Failure 500 {object} responses.Error "Something went wrong"

func (h *TodoHandler) Remove(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, errors.IdIsRequired)
	}

	err := h.TodoService.Remove(id)
	if err != nil {
		if err == errors.RecordNotExists {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errors.ErrorRemoving)
	}

	ctx.JSON(http.StatusOK, responses.Empty{Status: true})
}

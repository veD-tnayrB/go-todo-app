package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/veD-tnayrB/todo-app/common/models"
	"github.com/veD-tnayrB/todo-app/common/responses"
)

// @Summary Create todo
// @Description Allows you to create a todo by passing your todo structure
// @Router /todos [post]
// @Version 1.0
// @Tags todos
// @Param todo body models.Todo true "Todo details"
// @Success 200 {object} responses.Empty
// @Failure 400 {object} responses.Error "Malformed body"
// @Failure 400 {object} responses.Error "Title is required"
// @Failure 500 {object} responses.Error "Something went wrong"
func (h *TodoHandler) Create(ctx *gin.Context) {
	body := models.Todo{}
	h.Logger.Info("Handler: Error Handling CREATE request")
	if err := ctx.ShouldBindBodyWithJSON(body); err != nil {
		h.Logger.Warn("Handler: Couldnt parse body from JSON", "error", ErrMalformedBody)
		ctx.JSON(http.StatusBadRequest, ErrMalformedBody)
		return
	}

	if body.Title == "" {
		h.Logger.Warn("Handler: Title is required", "error", ErrTitleIsRequired)
		ctx.JSON(http.StatusBadRequest, ErrTitleIsRequired)
		return
	}

	err := h.TodoService.Create(&body)
	if err != nil {
		h.Logger.Warn("Handler: Error creating todo", "error", err, "body", body)
		ctx.JSON(http.StatusInternalServerError, responses.Error{Status: false, Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, responses.Empty{Status: true})
}

// @Summary Updating todo
// @Description Allows you to update a todo by passing its id and the todo structure
// @Router /todos/:id [put]
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
	h.Logger.Info("Handler: Error Handling UPDATE request")
	if err := ctx.ShouldBindBodyWithJSON(body); err != nil {
		h.Logger.Warn("Handler: Couldnt parse body from JSON", "error", ErrMalformedBody)
		ctx.JSON(http.StatusBadRequest, ErrMalformedBody)
		return
	}

	if body.Title == "" {
		h.Logger.Warn("Handler: Title is required", "error", ErrTitleIsRequired)
		ctx.JSON(http.StatusBadRequest, ErrTitleIsRequired)
		return
	}

	err := h.TodoService.Update(&body)
	if err != nil {
		if err == ErrRecordNotExists {
			h.Logger.Error("Handler: Error updating todo", "error", err, "body", body)
			ctx.JSON(http.StatusInternalServerError, responses.Error{Status: false, Message: err.Error(), Code: http.StatusInternalServerError})
			return
		}

		h.Logger.Error("Handler: Error updating todo", "error", err, "body", body)
		ctx.JSON(http.StatusInternalServerError, ErrErrorSaving)
		return
	}

	ctx.JSON(http.StatusOK, responses.Empty{Status: true})
}

// @Summary Removing todo
// @Description Allows you to remove a todo by passing its id
// @Router /todo/:id [delete]
// @Version 1.0
// @Tags todos
// @Param id identifier string true "Todo id"
// @Success 200 {object} responses.Empty
// @Failure 400 {object} responses.Error "Id is required"
// @Failure 500 {object} responses.Error "Record not exists"
// @Failure 500 {object} responses.Error "Something went wrong"

func (h *TodoHandler) Remove(ctx *gin.Context) {
	id := ctx.Param("id")
	h.Logger.Info("Handler: Error Handling REMOVE request", "id", id)
	if id == "" {
		h.Logger.Warn("Handler: id is required", "id", id, "error", ErrIdIsRequired)
		ctx.JSON(http.StatusBadRequest, ErrIdIsRequired)
		return
	}

	err := h.TodoService.Remove(id)
	if err != nil {
		if err == ErrRecordNotExists {
			h.Logger.Warn("Handler: record doesnt exists", "id", id)
			ctx.JSON(http.StatusNotFound, responses.Error{Status: false, Code: http.StatusNotFound, Message: err.Error()})
			return
		}

		h.Logger.Warn("Handler: Error trying to remove register", "id", id, "error", err)
		ctx.JSON(http.StatusInternalServerError, responses.Error{Status: false, Code: http.StatusInternalServerError, Message: ErrErrorRemoving.Error()})
		return
	}

	ctx.JSON(http.StatusOK, responses.Empty{Status: true})
}

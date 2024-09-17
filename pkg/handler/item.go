package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	todo "todo-app"
)

// @Summary Create Todo Item
// @Tags items
// @Description Create a new item in the todo list
// @ID create-item
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Param input body todo.TodoItem true "Todo Item Input"
// @Success 200 {object} map[string]interface{} "id"
// @Failure 400 {object} errorResponse "Invalid input body or list ID"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.Create(userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get All Todo Items
// @Tags items
// @Description Get all items from a specific todo list
// @ID get-all-items
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Success 200 {array} todo.TodoItem
// @Failure 400 {object} errorResponse "Invalid list ID"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /lists/{id}/items [get]
func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

// @Summary Get Todo Item By ID
// @Tags items
// @Description Get a specific item by its ID from a todo list
// @ID get-item-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {object} todo.TodoItem
// @Failure 400 {object} errorResponse "Invalid item ID"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /items/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

// @Summary Update Todo Item
// @Tags items
// @Description Update an existing todo item by its ID
// @ID update-item
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Param input body todo.UpdateItemInput true "Update Todo Item Input"
// @Success 200 {object} statusResponse "ok"
// @Failure 400 {object} errorResponse "Invalid item ID or input body"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /items/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete Todo Item
// @Tags items
// @Description Delete an existing todo item by its ID
// @ID delete-item
// @Accept  json
// @Produce  json
// @Param id path int true "Item ID"
// @Success 200 {object} statusResponse "ok"
// @Failure 400 {object} errorResponse "Invalid item ID"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")
		return
	}

	err = h.services.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

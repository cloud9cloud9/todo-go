package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	todo "todo-app"
)

// @Summary Create Todo List
// @Tags lists
// @Description Create a new todo list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body todo.TodoList true "Todo List Input"
// @Success 200 {object} map[string]interface{} "id of the created list"
// @Failure 400 {object} errorResponse "Invalid User ID or input body"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var intput todo.TodoList
	if err := c.BindJSON(&intput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, intput)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

// @Summary Get All Todo Lists
// @Tags lists
// @Description Get all todo lists for a user
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 400 {object} errorResponse "Invalid User ID"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// @Summary Get Todo List By ID
// @Tags lists
// @Description Get a specific todo list by its ID
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Success 200 {object} todo.TodoList
// @Failure 400 {object} errorResponse "Invalid list ID"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Update Todo List
// @Tags lists
// @Description Update an existing todo list by its ID
// @ID update-list
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Param input body todo.UpdateListInput true "Update Todo List Input"
// @Success 200 {object} statusResponse "ok"
// @Failure 400 {object} errorResponse "Invalid list ID or input body"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userId, listId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

// @Summary Delete Todo List
// @Tags lists
// @Description Delete an existing todo list by its ID
// @ID delete-list
// @Accept  json
// @Produce  json
// @Param id path int true "List ID"
// @Success 200 {object} statusResponse "ok"
// @Failure 400 {object} errorResponse "Invalid list ID"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

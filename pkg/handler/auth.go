package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	todo "todo-app"
)

type sighInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Sign In
// @Tags auth
// @Description User sign in
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param input body sighInInput true "Sign In Input"
// @Success 200 {object} map[string]interface{} "token"
// @Failure 400 {object} errorResponse "Invalid input body"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /auth/sign-in [post]

func (h *Handler) signIn(c *gin.Context) {
	var input sighInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

// @Summary Sign Up
// @Tags auth
// @Description User sign up
// @ID sign-up
// @Accept  json
// @Produce  json
// @Param input body todo.User true "Sign Up Input"
// @Success 200 {object} map[string]interface{} "id"
// @Failure 400 {object} errorResponse "Invalid input body"
// @Failure 500 {object} errorResponse "Internal Server Error"
// @Router /sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

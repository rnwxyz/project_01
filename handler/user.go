package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rnwxyz/project_01/helper"
	"github.com/rnwxyz/project_01/user"
)

type userHandler struct {
	userSarvice user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Akun gagal di buat", http.StatusUnprocessableEntity, "false", helper.APIValidation(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userSarvice.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Akun gagal di buat", http.StatusBadRequest, "false", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormaterUser(newUser, "token")

	response := helper.APIResponse("Akun berhasil di buat", http.StatusOK, "true", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Login gagal", http.StatusUnprocessableEntity, "false", helper.APIValidation(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userSarvice.Login(input)
	if err != nil {
		response := helper.APIResponse("Akun gagal di buat", http.StatusBadRequest, "false", helper.APIError(err))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormaterUser(loggedInUser, "tokentoken")

	response := helper.APIResponse("Login berhasil", http.StatusOK, "true", formatter)

	c.JSON(http.StatusOK, response)
}

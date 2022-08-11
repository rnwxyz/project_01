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
		response := helper.APIResponse("Akun gagal di buat", http.StatusBadRequest, "false", helper.APIError(err))
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

func (h *userHandler) CheckEmailAvailable(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "false", helper.APIValidation(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	EmailIsAvailable, err := h.userSarvice.EmailIsAvailable(input)
	if err != nil {
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "false", helper.APIValidation(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": EmailIsAvailable,
	}

	metaMessage := "Email sudah digunakan"

	if EmailIsAvailable {
		metaMessage = "Email bisa digunakan"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "true", data)

	c.JSON(http.StatusOK, response)
}

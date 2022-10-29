package controllers

import (
	"mygram/helpers"
	"mygram/params"
	"mygram/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (h *UserController) UserRegisterController(c *gin.Context) {
	userRequest := params.UserRegisterRequest{}
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}

	createUser, err := h.service.Register(userRequest)
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.InternalServerError(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.SuccessCreateResponse(createUser, "Register Success"))
}

func (h *UserController) UserLoginController(c *gin.Context) {
	loginRequest := params.UserLoginRequest{}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}

	token, err := h.service.Login(loginRequest)
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.InternalServerError(err))
		return
	}

	// signedToken := token.token
	// c.SetCookie("token", signedToken, 3600, "", "", false, true)
	
	helpers.WriteJsonRespnse(c, helpers.SuccessResponse(token,"Login Success"))
}

func (h *UserController) UserUpdateController(c *gin.Context) {
	user := params.UserUpdateRequest{}
	if err := c.ShouldBindJSON(&user); err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}

	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}

	updatedUser, err := h.service.Update(uint(userId), user)
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}

	helpers.WriteJsonRespnse(c, helpers.SuccessResponse(updatedUser, "Update Success"))
}

func (h *UserController) DeleteUserController(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}
	err = h.service.Delete(uint(userId))
	if err != nil {
		helpers.WriteJsonRespnse(c, helpers.BadRequestResponse(err))
		return
	}

	// c.SetCookie("token", "", 0, "", "", false, true)

	helpers.WriteJsonRespnse(c, helpers.DeleteSuccess("Your account has been successfully deleted"))
}
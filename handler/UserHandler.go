package handler

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg/errs"
	"esmartcare/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

// ==================================================

func (u *UserHandler) Register(ctx *gin.Context) {
	var requestBody dto.RegisterRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	registeredUser, err := u.userService.Register(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, registeredUser)
}

func (u *UserHandler) Login(ctx *gin.Context) {
	var requestBody dto.LoginRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	token, err := u.userService.Login(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (u *UserHandler) GettAllUsers(ctx *gin.Context) {
	var AllUsers []dto.GetAllUsersResponse
	jenisAkun := ctx.Query("jenis-akun")

	AllUsers, err := u.userService.GetAllUsers(jenisAkun)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, AllUsers)
}

func (u *UserHandler) GettAllUsersNotValidate(ctx *gin.Context) {
	var AllUsers []dto.GetAllUsersResponse
	jenisAkun := ctx.Query("jenis-akun")

	AllUsers, err := u.userService.GetAllUsersNotValidate(jenisAkun)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, AllUsers)
}

func (u *UserHandler) UpdateUser(ctx *gin.Context) {

	userEmail := ctx.Param("email")
	updatedUser, err := u.userService.UpdateUser(userEmail)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, updatedUser)
}

func (u *UserHandler) DeleteUser(ctx *gin.Context) {

	payload := entity.User{}
	userEmail := ctx.Param("email")

	payload.Email = userEmail

	response, err := u.userService.DeleteUser(&payload)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

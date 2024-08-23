package handler

import (
	"encoding/json"
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg/errs"
	"esmartcare/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	AdminService service.AdminService
}

func NewAdminHandler(AdminService service.AdminService) *AdminHandler {
	return &AdminHandler{AdminService}
}

func (s *AdminHandler) CreateAdmin(ctx *gin.Context) {
	var requestBody dto.CreateAdminRequest

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	// Retrieve JSON part from form-data
	jsonData := ctx.PostForm("Admin")
	if jsonData != "" {
		if err := json.Unmarshal([]byte(jsonData), &requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	registeredUser, err2 := s.AdminService.CreateOrUpdateAdminWithProfilPhoto(userData.Email, &requestBody, ctx)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusCreated, registeredUser)
}

func (s *AdminHandler) UploadProfileImage(ctx *gin.Context) {

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	registeredUser, err2 := s.AdminService.UpdateProfilPhoto(userData.Email, ctx)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}
	ctx.JSON(http.StatusAccepted, registeredUser)
}

func (s *AdminHandler) CreateOrUpdateAdmin(ctx *gin.Context) {
	var requestBody dto.CreateAdminRequest

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	// Retrieve JSON part from form-data

	registeredUser, err2 := s.AdminService.CreateOrUpdateAdmin(userData.Email, &requestBody)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusCreated, registeredUser)
}

func (s *AdminHandler) GetAdmin(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	// Retrieve JSON part from form-data

	admin, err := s.AdminService.GetAdmin(userData.Email)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, admin)
}

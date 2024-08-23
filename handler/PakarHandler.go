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

type PakarHandler struct {
	PakarService service.PakarService
}

func NewPakarHandler(PakarService service.PakarService) *PakarHandler {
	return &PakarHandler{PakarService}
}

func (s *PakarHandler) CreatePakar(ctx *gin.Context) {
	var requestBody dto.CreatePakarRequest

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	// Retrieve JSON part from form-data
	jsonData := ctx.PostForm("Pakar")
	if jsonData != "" {
		if err := json.Unmarshal([]byte(jsonData), &requestBody); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	registeredUser, err2 := s.PakarService.CreateOrUpdatePakarWithProfilPhoto(userData.Email, &requestBody, ctx)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusCreated, registeredUser)
}

func (s *PakarHandler) UploadProfileImage(ctx *gin.Context) {

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	registeredUser, err2 := s.PakarService.UpdateProfilPhoto(userData.Email, ctx)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}
	ctx.JSON(http.StatusAccepted, registeredUser)
}

func (s *PakarHandler) CreateOrUpdatePakar(ctx *gin.Context) {
	var requestBody dto.CreatePakarRequest

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

	registeredUser, err2 := s.PakarService.CreateOrUpdatePakar(userData.Email, &requestBody)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusCreated, registeredUser)
}

func (s *PakarHandler) GetPakar(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	// Retrieve JSON part from form-data

	pakar, err := s.PakarService.GetPakar(userData.Email)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusOK, pakar)
}

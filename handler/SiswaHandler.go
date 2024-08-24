package handler

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg/errs"
	"esmartcare/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SiswaHandler struct {
	siswaService service.SiswaService
}

func NewSiswaHandler(siswaService service.SiswaService) *SiswaHandler {
	return &SiswaHandler{siswaService}
}

func (s *SiswaHandler) UploadProfileImage(ctx *gin.Context) {

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	registeredUser, err2 := s.siswaService.UpdateProfilPhoto(userData.Email, ctx)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}
	ctx.JSON(http.StatusAccepted, registeredUser)
}

func (s *SiswaHandler) CreateOrUpdateSiswa(ctx *gin.Context) {
	var requestBody dto.CreateSiswaRequest

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

	registeredUser, err2 := s.siswaService.CreateOrUpdateSiswa(userData.Email, &requestBody)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusCreated, registeredUser)
}

func (h *SiswaHandler) GetAllSiswaWithPemeriksaan(c *gin.Context) {
	siswa, err := h.siswaService.GetAllSiswaWithPemeriksaan()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, siswa)
}

func (h *SiswaHandler) GetSiswa(ctx *gin.Context) {

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	siswa, err := h.siswaService.GetSiswa(userData.Email)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}
	ctx.JSON(http.StatusOK, siswa)
}

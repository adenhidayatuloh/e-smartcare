package handler

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg/errs"
	"esmartcare/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PemeriksaanHandler struct {
	service service.PemeriksaanService
}

func NewPemeriksaanHandler(service service.PemeriksaanService) *PemeriksaanHandler {
	return &PemeriksaanHandler{service: service}
}

func (h *PemeriksaanHandler) GetAllPemeriksaan(ctx *gin.Context) {
	pemeriksaans, err := h.service.GetAllPemeriksaan()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, pemeriksaans)
}

func (h *PemeriksaanHandler) CreatePemeriksaan(ctx *gin.Context) {

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}
	var request dto.CreateUpdatePemeriksaanRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	request.Email = userData.Email

	pemeriksaan, err := h.service.CreatePemeriksaan(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{

		"email":      pemeriksaan.Email,
		"waktu":      pemeriksaan.Waktu,
		"foto":       pemeriksaan.Foto,
		"tinggi":     pemeriksaan.Tinggi,
		"berat":      pemeriksaan.Berat,
		"keterangan": pemeriksaan.Keterangan,
	})
}

func (h *PemeriksaanHandler) GetPemeriksaanByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	pemeriksaans, err := h.service.GetPemeriksaanByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, pemeriksaans)
}

func (h *PemeriksaanHandler) DeletePemeriksaanByEmail(ctx *gin.Context) {
	email := ctx.Param("email")

	if err := h.service.DeletePemeriksaanByEmail(email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (s *PemeriksaanHandler) UploadPhotoPemeriksaan(ctx *gin.Context) {

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	registeredUser, err2 := s.service.UpdatePhotoPemeriksaan(userData.Email, ctx)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{
		"email": registeredUser.Email,
		"foto":  registeredUser.Foto,
	})
}

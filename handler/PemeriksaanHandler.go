package handler

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg/errs"
	"esmartcare/service"
	"net/http"
	"strconv"
	"strings"

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

	tinggiStr := ctx.PostForm("tinggi")
	beratStr := ctx.PostForm("berat")
	keterangan := ctx.PostForm("keterangan")

	keterangan = strings.ToLower(keterangan)

	tinggi, err := strconv.ParseFloat(tinggiStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input for tinggi"})
		return
	}

	berat, err := strconv.ParseFloat(beratStr, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input for berat"})
		return
	}

	request.Email = userData.Email
	request.Berat = berat
	request.Tinggi = tinggi
	request.Keterangan = keterangan

	pemeriksaan, err := h.service.CreatePemeriksaan(request, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         pemeriksaan.IdPemeriksaan,
		"email":      pemeriksaan.Email,
		"waktu":      pemeriksaan.Waktu,
		"foto":       pemeriksaan.Foto,
		"tinggi":     pemeriksaan.Tinggi,
		"berat":      pemeriksaan.Berat,
		"keterangan": pemeriksaan.Keterangan,
	})
}

func (h *PemeriksaanHandler) GetPemeriksaanByEmail(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}
	keterangan := ctx.Query("keterangan")
	pemeriksaans, err := h.service.GetPemeriksaanByEmail(userData.Email, keterangan)
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

func (h *PemeriksaanHandler) DeletePemeriksaanById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeletePemeriksaanById(id); err != nil {
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

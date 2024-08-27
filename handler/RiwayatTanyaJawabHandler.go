package handler

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg/errs"
	"esmartcare/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RiwayatTanyaJawabHandler struct {
	service service.RiwayatTanyaJawabService
}

func NewRiwayatTanyaJawabHandler(service service.RiwayatTanyaJawabService) *RiwayatTanyaJawabHandler {
	return &RiwayatTanyaJawabHandler{service: service}
}

func (h *RiwayatTanyaJawabHandler) GetRiwayatByEmail(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}
	riwayat, err := h.service.GetRiwayatByEmail(userData.Email)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, riwayat)
}

func (h *RiwayatTanyaJawabHandler) GetAllDataRiwayat(ctx *gin.Context) {

	riwayat, err := h.service.GetAllDataRiwayat()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, riwayat)
}

func (h *RiwayatTanyaJawabHandler) CreateRiwayat(ctx *gin.Context) {
	var request dto.CreateUpdateRiwayatTanyaJawabRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	request.Email = userData.Email

	riwayat, err := h.service.CreateRiwayat(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{

		"id":         riwayat.Id,
		"email":      riwayat.Email,
		"waktu":      riwayat.Waktu,
		"pertanyaan": riwayat.Pertanyaan,
		"jawaban":    riwayat.Jawaban,
	})
}

func (h *RiwayatTanyaJawabHandler) DeleteRiwayatById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteRiwayatById(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}

func (h *RiwayatTanyaJawabHandler) DeleteRiwayatByEmail(ctx *gin.Context) {
	emailParam := ctx.Param("email")

	if err := h.service.DeleteRiwayatByEmail(emailParam); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Record with email " + emailParam + " deleted"})
}

func (h *RiwayatTanyaJawabHandler) DeleteAllRiwayatByUserLogin(ctx *gin.Context) {

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	if err := h.service.DeleteRiwayatByEmail(userData.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Record with email " + userData.Email + " deleted"})
}

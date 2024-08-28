package handler

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg"
	"esmartcare/pkg/errs"
	"esmartcare/service"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/blevesearch/bleve/v2"
	"github.com/gin-gonic/gin"
)

type TanyaJawabHandler struct {
	service service.TanyaJawabService
}

func NewTanyaJawabHandler(service service.TanyaJawabService) *TanyaJawabHandler {
	return &TanyaJawabHandler{service: service}
}

func (h *TanyaJawabHandler) GetTanyaJawab(ctx *gin.Context) {
	isvalidate := ctx.Query("isvalidate")

	var tanyaJawab []entity.TanyaJawab
	var err error

	if isvalidate == "true" {
		tanyaJawab, err = h.service.GetTanyaJawabByValidationStatus(true)
	} else if isvalidate == "false" {
		tanyaJawab, err = h.service.GetTanyaJawabByValidationStatus(false)
	} else {
		tanyaJawab, err = h.service.GetAllTanyaJawab()
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tanyaJawab)
}

func (h *TanyaJawabHandler) CreateTanyaJawab(ctx *gin.Context) {
	var request dto.CreateUpdateTanyaJawabRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tanyaJawab, err := h.service.CreateTanyaJawab(request)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, tanyaJawab)
}

func (h *TanyaJawabHandler) UpdateTanyaJawab(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var request dto.CreateUpdateTanyaJawabRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	tanyaJawab, err2 := h.service.UpdateTanyaJawab(id, &request)
	if err2 != nil {
		ctx.JSON(err2.StatusCode(), err2)
		return
	}

	ctx.JSON(http.StatusOK, tanyaJawab)
}

func (h *TanyaJawabHandler) UpdateValidator(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	tanyaJawab, err := h.service.UpdateValidator(id, userData.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tanyaJawab)
}

func (h *TanyaJawabHandler) DeleteTanyaJawab(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.DeleteTanyaJawab(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Record deleted"})
}

func (h *TanyaJawabHandler) ChatSimmilarityBot(ctx *gin.Context) {
	var request dto.ChatbotSimillarityRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tanyaJawab, err := h.service.GetSimillaryQuestion(request.Pertanyaan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": tanyaJawab})
}

func (h *TanyaJawabHandler) ChatBot(ctx *gin.Context) {
	var request dto.ChatbotSimillarityRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Jawaban, Kemiripan, err := h.service.GetChatQuestion(request.Pertanyaan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"pertanyaan": request.Pertanyaan,
		"jawaban":    Jawaban,
		"kemiripan":  Kemiripan})
}

func (h *TanyaJawabHandler) Update_Bot(c *gin.Context) {

	var faqs []entity.FAQ

	faqs, err := h.service.GetChatBotUpdate()

	if err == nil {

		// Buat atau buka indeks Bleve
		indexMapping := bleve.NewIndexMapping()
		index, err := bleve.Open("faq.bleve")
		if err == bleve.ErrorIndexPathDoesNotExist {
			index, err = pkg.CreateIndex("faq.bleve", indexMapping)
			if err != nil {
				log.Fatal(err)
			}
		} else if err != nil {
			log.Fatal(err)
		}
		defer index.Close()

		// Hapus semua dokumen dari indeks sebelum pengindeksan ulang
		if err := pkg.DeleteAllDocuments(index); err != nil {
			log.Fatal(err)
		}

		// Indexing data
		for i, faq := range faqs {
			err := index.Index(fmt.Sprintf("%d", i), faq)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "sukses update bot",
	})
}

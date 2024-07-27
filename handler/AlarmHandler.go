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

type AlarmHandler struct {
	alarmService service.AlarmService
}

func NewAlarmHandler(as service.AlarmService) *AlarmHandler {
	return &AlarmHandler{as}
}

func (h *AlarmHandler) GetAllAlarms(c *gin.Context) {
	alarms, err := h.alarmService.GetAllAlarms()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alarms)
}

func (h *AlarmHandler) GetAlarmsByEmail(ctx *gin.Context) {

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	alarms, err := h.alarmService.GetAlarmsByEmail(userData.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, alarms)
}

func (h *AlarmHandler) CreateAlarm(ctx *gin.Context) {
	var alarmDTO dto.CreateUpdateAlarmRequestResponse

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	if err := ctx.ShouldBindJSON(&alarmDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, errs.NewUnprocessableEntity(err.Error()))
		return
	}

	alarmDTO.Email = userData.Email

	if err := h.alarmService.CreateAlarm(&alarmDTO); err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}
	ctx.JSON(http.StatusCreated, alarmDTO)
}

func (h *AlarmHandler) DeleteAlarmByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errs.NewUnprocessableEntity("Invalid ID"))
		return
	}
	if err := h.alarmService.DeleteAlarmByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Alarm deleted"})
}

func (h *AlarmHandler) UpdateAlarm(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errs.NewUnprocessableEntity("Invalid ID"))
		return
	}

	var alarmDTO dto.CreateUpdateAlarmRequestResponse

	userData, ok := ctx.MustGet("userData").(*entity.User)

	if !ok {
		newError := errs.NewBadRequest("Failed to get user data")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	if err := ctx.ShouldBindJSON(&alarmDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, errs.NewUnprocessableEntity(err.Error()))
		return
	}

	alarmDTO.Email = userData.Email

	Updated, err := h.alarmService.UpdateAlarm(id, &alarmDTO)

	alarmResponse := dto.CreateUpdateAlarmRequestResponse{

		Email:        Updated.Email,
		Keterangan:   Updated.Keterangan,
		TanggalMulai: Updated.TanggalMulai,
		Jam:          Updated.Jam,
		Pengulangan:  Updated.Pengulangan,
		Status:       Updated.Status,
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, alarmResponse)
}

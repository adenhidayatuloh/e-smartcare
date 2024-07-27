package alarmmysql

import (
	"esmartcare/entity"
	"esmartcare/pkg/errs"
	"fmt"
	"log"

	"gorm.io/gorm"

	AlarmRepository "esmartcare/repository/alarmRepository"
)

type alarmRepository struct {
	db *gorm.DB
}

func NewAlarmRepository(db *gorm.DB) AlarmRepository.AlarmRepository {
	return &alarmRepository{db}
}

func (r *alarmRepository) GetAllAlarms() ([]entity.Alarm, errs.MessageErr) {
	var alarms []entity.Alarm
	err := r.db.Order("email ASC, tanggal_mulai ASC").Find(&alarms).Error

	if err != nil {
		return nil, errs.NewNotFound("Data not found")
	}
	return alarms, nil
}

func (r *alarmRepository) GetAlarmsByEmail(email string) ([]entity.Alarm, errs.MessageErr) {
	var alarms []entity.Alarm
	err := r.db.Where("email = ?", email).Order("tanggal_mulai ASC").Find(&alarms).Error

	if err != nil {
		return nil, errs.NewNotFound("Alarm with email " + email + " not found")
	}
	return alarms, nil
}

func (r *alarmRepository) CreateAlarm(alarm *entity.Alarm) errs.MessageErr {

	err := r.db.Create(alarm).Error

	if err != nil {
		return errs.NewBadRequest("Cannot create alarm")
	}
	return nil
}

func (r *alarmRepository) DeleteAlarmByID(id int) errs.MessageErr {

	err := r.db.Delete(&entity.Alarm{}, id).Error

	if err != nil {
		return errs.NewBadRequest(fmt.Sprintf("Cannot delete alarm with id %d", id))
	}
	return nil
}

func (r *alarmRepository) UpdateAlarm(oldAlarm *entity.Alarm, newAlarm *entity.Alarm) (*entity.Alarm, errs.MessageErr) {
	if err := r.db.Model(oldAlarm).Updates(newAlarm).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewUnprocessableEntity(fmt.Sprintf("Failed to update alarm %s", oldAlarm.Email))
	}

	return oldAlarm, nil
}

// GetAlarmById implements alarmrepository.AlarmRepository.
func (r *alarmRepository) GetAlarmById(id int) (*entity.Alarm, errs.MessageErr) {
	var alarm entity.Alarm

	if err := r.db.First(&alarm, id).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("Alarm with id %d is not found", id))
	}

	return &alarm, nil
}

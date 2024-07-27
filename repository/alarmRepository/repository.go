package alarmrepository

import (
	"esmartcare/entity"
	"esmartcare/pkg/errs"
)

type AlarmRepository interface {
	GetAllAlarms() ([]entity.Alarm, errs.MessageErr)
	GetAlarmsByEmail(email string) ([]entity.Alarm, errs.MessageErr)
	GetAlarmById(id int) (*entity.Alarm, errs.MessageErr)
	CreateAlarm(alarm *entity.Alarm) errs.MessageErr
	DeleteAlarmByID(id int) errs.MessageErr
	UpdateAlarm(oldAlarm *entity.Alarm, newAlarm *entity.Alarm) (*entity.Alarm, errs.MessageErr)
}

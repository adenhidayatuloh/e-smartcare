package service

import (
	"esmartcare/dto"
	"esmartcare/entity"
	"esmartcare/pkg"
	"esmartcare/pkg/errs"

	AlarmRepository "esmartcare/repository/alarmRepository"
)

type AlarmService interface {
	GetAllAlarms() ([]entity.Alarm, errs.MessageErr)
	GetAlarmsByEmail(email string) ([]entity.Alarm, errs.MessageErr)
	CreateAlarm(alarm *dto.CreateUpdateAlarmRequestResponse) errs.MessageErr
	DeleteAlarmByID(id int) errs.MessageErr
	UpdateAlarm(id int, newAlarm *dto.CreateUpdateAlarmRequestResponse) (*entity.Alarm, errs.MessageErr)
}

type alarmService struct {
	alarmRepo AlarmRepository.AlarmRepository
}

func NewAlarmService(ar AlarmRepository.AlarmRepository) AlarmService {
	return &alarmService{ar}
}

func (s *alarmService) GetAllAlarms() ([]entity.Alarm, errs.MessageErr) {
	return s.alarmRepo.GetAllAlarms()
}

func (s *alarmService) GetAlarmsByEmail(email string) ([]entity.Alarm, errs.MessageErr) {
	return s.alarmRepo.GetAlarmsByEmail(email)
}

func (s *alarmService) CreateAlarm(alarmDTO *dto.CreateUpdateAlarmRequestResponse) errs.MessageErr {

	err := pkg.ValidateStruct(alarmDTO)
	if err != nil {
		return err
	}

	alarm := entity.Alarm{
		Email:        alarmDTO.Email,
		Keterangan:   alarmDTO.Keterangan,
		TanggalMulai: alarmDTO.TanggalMulai,
		Jam:          alarmDTO.Jam,
		Pengulangan:  alarmDTO.Pengulangan,
		Status:       alarmDTO.Status,
	}

	return s.alarmRepo.CreateAlarm(&alarm)
}

func (s *alarmService) DeleteAlarmByID(id int) errs.MessageErr {
	return s.alarmRepo.DeleteAlarmByID(id)
}

func (s *alarmService) UpdateAlarm(id int, alarmDTO *dto.CreateUpdateAlarmRequestResponse) (*entity.Alarm, errs.MessageErr) {

	oldAlarm, errGet := s.alarmRepo.GetAlarmById(id)

	if errGet != nil {
		return nil, errGet
	}

	err := pkg.ValidateStruct(alarmDTO)
	if err != nil {
		return nil, err
	}

	alarm := entity.Alarm{
		ID:           id,
		Email:        alarmDTO.Email,
		Keterangan:   alarmDTO.Keterangan,
		TanggalMulai: alarmDTO.TanggalMulai,
		Jam:          alarmDTO.Jam,
		Pengulangan:  alarmDTO.Pengulangan,
		Status:       alarmDTO.Status,
	}

	return s.alarmRepo.UpdateAlarm(oldAlarm, &alarm)

}

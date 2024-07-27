package pakarMysql

import (
	"errors"
	"esmartcare/entity"
	"esmartcare/pkg/errs"
	Pakarrepository "esmartcare/repository/pakarRepository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type pakarMysql struct {
	db *gorm.DB
}

// UpdatePakar implements Pakarrepository.PakarRepository.
func (s *pakarMysql) UpdatePakar(oldPakar *entity.Pakar, newPakar *entity.Pakar) (*entity.Pakar, errs.MessageErr) {
	if err := s.db.Model(oldPakar).Updates(newPakar).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewUnprocessableEntity(fmt.Sprintf("Failed to update email %s", oldPakar.Email))
	}

	return oldPakar, nil
}

// CreatePakar implements Pakarrepository.PakarRepository.
func (s *pakarMysql) CreatePakar(Pakar *entity.Pakar) (*entity.Pakar, errs.MessageErr) {

	if err := s.db.Create(Pakar).Error; err != nil {
		log.Println("Error:", err.Error())

		return nil, errs.NewInternalServerError("Failed to create Pakar ")
	}

	return Pakar, nil
}

func NewpakarMysql(db *gorm.DB) Pakarrepository.PakarRepository {
	return &pakarMysql{db}
}

// GetPakarByEmail implements Pakarrepository.PakarRepository.
func (s *pakarMysql) GetPakarByEmail(email string) (*entity.Pakar, errs.MessageErr) {
	var Pakar entity.Pakar

	if err := s.db.First(&Pakar, "email = ?", email).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle record not found error...

			return nil, errs.NewNotFound(fmt.Sprintf("Pakar with email %s is not found", email))
		}
		return nil, errs.NewBadRequest("cannot get Pakar")
	}

	return &Pakar, nil
}

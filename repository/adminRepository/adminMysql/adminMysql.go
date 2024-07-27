package adminmysql

import (
	"errors"
	"esmartcare/entity"
	"esmartcare/pkg/errs"
	adminrepository "esmartcare/repository/adminRepository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type adminMySql struct {
	db *gorm.DB
}

// UpdateAdmin implements Adminrepository.AdminRepository.
func (s *adminMySql) UpdateAdmin(oldAdmin *entity.Admin, newAdmin *entity.Admin) (*entity.Admin, errs.MessageErr) {
	if err := s.db.Model(oldAdmin).Updates(newAdmin).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewUnprocessableEntity(fmt.Sprintf("Failed to update email %s", oldAdmin.Email))
	}

	return oldAdmin, nil
}

// CreateAdmin implements Adminrepository.AdminRepository.
func (s *adminMySql) CreateAdmin(admin *entity.Admin) (*entity.Admin, errs.MessageErr) {

	if err := s.db.Create(admin).Error; err != nil {
		log.Println("Error:", err.Error())

		return nil, errs.NewInternalServerError("Failed to create Admin ")
	}

	return admin, nil
}

func NewAdminMySql(db *gorm.DB) adminrepository.AdminRepository {
	return &adminMySql{db}
}

// GetAdminByEmail implements Adminrepository.AdminRepository.
func (s *adminMySql) GetAdminByEmail(email string) (*entity.Admin, errs.MessageErr) {
	var Admin entity.Admin

	if err := s.db.First(&Admin, "email = ?", email).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle record not found error...

			return nil, errs.NewNotFound(fmt.Sprintf("Admin with email %s is not found", email))
		}
		return nil, errs.NewBadRequest("cannot get Admin")
	}

	return &Admin, nil
}

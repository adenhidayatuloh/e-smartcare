package pemeriksaanmysql

import (
	"esmartcare/entity"
	"esmartcare/pkg/errs"
	PemeriksaanRepository "esmartcare/repository/pemeriksaanRepository"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type pemeriksaanRepository struct {
	db *gorm.DB
}

func NewPemeriksaanRepository(db *gorm.DB) PemeriksaanRepository.PemeriksaanRepository {
	return &pemeriksaanRepository{db: db}
}

func (r *pemeriksaanRepository) FindAll() ([]entity.Pemeriksaan, error) {
	var pemeriksaans []entity.Pemeriksaan
	if err := r.db.Preload("Siswa").Order("email asc, waktu asc").Find(&pemeriksaans).Error; err != nil {
		return nil, err
	}
	return pemeriksaans, nil
}

func (r *pemeriksaanRepository) Create(pemeriksaan entity.Pemeriksaan) (*entity.Pemeriksaan, error) {
	if err := r.db.Create(&pemeriksaan).Error; err != nil {
		return nil, err
	}
	return &pemeriksaan, nil
}

func (r *pemeriksaanRepository) FindByEmail(email string) ([]entity.Pemeriksaan, error) {
	var pemeriksaans []entity.Pemeriksaan
	if err := r.db.Where("email = ?", email).Order("waktu asc").Find(&pemeriksaans).Error; err != nil {
		return nil, err
	}
	return pemeriksaans, nil
}

func (r *pemeriksaanRepository) DeleteByEmail(email string) error {
	if err := r.db.Where("email = ?", email).Delete(&entity.Pemeriksaan{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *pemeriksaanRepository) GetPemeriksaanByEmail(email string) (*entity.Pemeriksaan, error) {
	var pemeriksaan entity.Pemeriksaan

	if err := r.db.First(&pemeriksaan, "email = ?", email).Error; err != nil {
		return nil, errs.NewNotFound(fmt.Sprintf("User with email %s is not found", email))
	}

	return &pemeriksaan, nil
}

func (s *pemeriksaanRepository) UpdatePemeriksaan(oldPemeriksaan *entity.Pemeriksaan, newPemeriksaan *entity.Pemeriksaan) (*entity.Pemeriksaan, errs.MessageErr) {
	if err := s.db.Model(oldPemeriksaan).Updates(newPemeriksaan).Error; err != nil {
		log.Println("Error:", err.Error())
		return nil, errs.NewUnprocessableEntity(fmt.Sprintf("Failed to update pemeriksaan email %s", oldPemeriksaan.Email))
	}

	return oldPemeriksaan, nil
}

package siswamysql

import (
	"errors"
	"esmartcare/entity"
	"esmartcare/pkg/errs"
	siswarepository "esmartcare/repository/siswaRepository"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type siswaMySql struct {
	db *gorm.DB
}

// UpdateSiswa implements siswarepository.SiswaRepository.
func (s *siswaMySql) UpdateSiswa(oldSiswa *entity.Siswa, newSiswa *entity.Siswa) (*entity.Siswa, errs.MessageErr) {
	if err := s.db.Model(oldSiswa).Updates(newSiswa).Error; err != nil {
		log.Println("Error:", err.Error())

		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1292:
				return nil, errs.NewUnprocessableEntity(fmt.Sprintf("Date not valid %s", newSiswa.TanggalLahir))
			default:
				return nil, errs.NewUnprocessableEntity(fmt.Sprintf("Failed to update email %s", oldSiswa.Email))
			}
		}
		return nil, errs.NewUnprocessableEntity(fmt.Sprintf("Failed to update email %s", oldSiswa.Email))
	}

	return oldSiswa, nil
}

// CreateSiswa implements siswarepository.SiswaRepository.
func (s *siswaMySql) CreateSiswa(siswa *entity.Siswa) (*entity.Siswa, errs.MessageErr) {

	if err := s.db.Create(siswa).Error; err != nil {
		log.Println("Error:", err.Error())

		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1292:
				return nil, errs.NewUnprocessableEntity(fmt.Sprintf("Date not valid %s", siswa.TanggalLahir))
			default:
				return nil, errs.NewInternalServerError("Failed to create siswa ")
			}
		}
		return nil, errs.NewInternalServerError("Failed to create siswa ")
	}

	return siswa, nil
}

func NewSiswaMySql(db *gorm.DB) siswarepository.SiswaRepository {
	return &siswaMySql{db}
}

// GetSiswaByEmail implements siswarepository.SiswaRepository.
func (s *siswaMySql) GetSiswaByEmail(email string) (*entity.Siswa, errs.MessageErr) {
	var siswa entity.Siswa

	if err := s.db.First(&siswa, "email = ?", email).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Handle record not found error...

			return nil, errs.NewNotFound(fmt.Sprintf("Siswa with email %s is not found", email))
		}
		return nil, errs.NewBadRequest("cannot get siswa")
	}

	return &siswa, nil
}

func (r *siswaMySql) GetAllSiswaWithPemeriksaan(keterangan string) ([]entity.Siswa_pemeriksaan, errs.MessageErr) {
	var siswa []entity.Siswa_pemeriksaan
	// err := r.db.Preload("Pemeriksaan").Order("email ASC , Pemeriksaan.waktu ASC").Find(&siswa).Error

	if keterangan != "normal" && keterangan != "gemuk" && keterangan != "stunting" {
		err := r.db.Preload("Pemeriksaan", func(db *gorm.DB) *gorm.DB {
			return db.Order("waktu ASC")
		}).Order("email ASC").Find(&siswa).Error

		if err != nil {
			return nil, errs.NewUnprocessableEntity("Cannot Get Siswa")
		}

		return siswa, nil

	}
	err := r.db.Preload("Pemeriksaan", func(db *gorm.DB) *gorm.DB {
		return db.Order("waktu ASC").Where("keterangan = ?", keterangan)
	}).Order("email ASC").Find(&siswa).Error

	if err != nil {
		return nil, errs.NewUnprocessableEntity("Cannot Get Siswa")
	}
	return siswa, nil
}

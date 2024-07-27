package riwayatmysql

import (
	"esmartcare/entity"

	"gorm.io/gorm"

	RiwayatTanyaJawabRepository "esmartcare/repository/riwayatTanyaJawabRepository"
)

type riwayatTanyaJawabRepository struct {
	db *gorm.DB
}

func NewRiwayatTanyaJawabRepository(db *gorm.DB) RiwayatTanyaJawabRepository.RiwayatTanyaJawabRepository {
	return &riwayatTanyaJawabRepository{db: db}
}

func (r *riwayatTanyaJawabRepository) FindByEmail(email string) ([]entity.RiwayatTanyaJawab, error) {
	var riwayat []entity.RiwayatTanyaJawab
	if err := r.db.Where("email = ?", email).Order("waktu asc").Find(&riwayat).Error; err != nil {
		return nil, err
	}
	return riwayat, nil
}

func (r *riwayatTanyaJawabRepository) FindById(id uint) (*entity.RiwayatTanyaJawab, error) {
	var riwayat entity.RiwayatTanyaJawab
	if err := r.db.Where("id = ?", id).Find(&riwayat).Error; err != nil {
		return nil, err
	}
	return &riwayat, nil
}

func (r *riwayatTanyaJawabRepository) Create(riwayat entity.RiwayatTanyaJawab) (entity.RiwayatTanyaJawab, error) {
	if err := r.db.Create(&riwayat).Error; err != nil {
		return entity.RiwayatTanyaJawab{}, err
	}
	return riwayat, nil
}

func (r *riwayatTanyaJawabRepository) DeleteById(id uint) error {
	if err := r.db.Delete(&entity.RiwayatTanyaJawab{}, id).Error; err != nil {
		return err
	}
	return nil
}

package tanyajawabmysql

import (
	"esmartcare/entity"

	TanyaJawabRepository "esmartcare/repository/tanyaJawabRepository"

	"gorm.io/gorm"
)

type tanyaJawabRepository struct {
	db *gorm.DB
}

func NewTanyaJawabRepository(db *gorm.DB) TanyaJawabRepository.TanyaJawabRepository {
	return &tanyaJawabRepository{db: db}
}

func (r *tanyaJawabRepository) FindAll() ([]entity.TanyaJawab, error) {
	var tanyaJawab []entity.TanyaJawab
	if err := r.db.Find(&tanyaJawab).Error; err != nil {
		return nil, err
	}
	return tanyaJawab, nil
}

func (r *tanyaJawabRepository) FindByValidationStatus(isValidated bool) ([]entity.TanyaJawab, error) {
	var tanyaJawab []entity.TanyaJawab
	query := r.db
	if isValidated {
		query = query.Where("validator IS NOT NULL AND validator != ''")
	} else {
		query = query.Where("validator IS NULL OR validator = ''")
	}
	if err := query.Find(&tanyaJawab).Error; err != nil {
		return nil, err
	}
	return tanyaJawab, nil
}

func (r *tanyaJawabRepository) FindByID(id int) (entity.TanyaJawab, error) {
	var tanyaJawab entity.TanyaJawab
	if err := r.db.First(&tanyaJawab, id).Error; err != nil {
		return entity.TanyaJawab{}, err
	}
	return tanyaJawab, nil
}

func (r *tanyaJawabRepository) Create(tanyaJawab entity.TanyaJawab) (entity.TanyaJawab, error) {
	if err := r.db.Create(&tanyaJawab).Error; err != nil {
		return entity.TanyaJawab{}, err
	}
	return tanyaJawab, nil
}

func (r *tanyaJawabRepository) Update(tanyaJawab entity.TanyaJawab) (entity.TanyaJawab, error) {
	if err := r.db.Save(&tanyaJawab).Error; err != nil {
		return entity.TanyaJawab{}, err
	}
	return tanyaJawab, nil
}

func (r *tanyaJawabRepository) Delete(id int) error {
	if err := r.db.Delete(&entity.TanyaJawab{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *tanyaJawabRepository) Get(id int) error {
	if err := r.db.Delete(&entity.TanyaJawab{}, id).Error; err != nil {
		return err
	}
	return nil
}

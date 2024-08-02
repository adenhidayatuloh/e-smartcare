package tanyajawabmysql

import (
	"esmartcare/entity"
	"log"

	TanyaJawabRepository "esmartcare/repository/tanyaJawabRepository"

	"gorm.io/gorm"
)

type tanyaJawabRepository struct {
	db *gorm.DB
}

// FindForChatbot implements TanyaJawabrepository.TanyaJawabRepository.
func (r *tanyaJawabRepository) FindForChatbot() ([]entity.FAQ, error) {
	// Raw SQL
	rows, err := r.db.Raw("SELECT pertanyaan, jawaban, validator FROM tanya_jawab WHERE validator IS NOT NULL AND validator != ''").Rows()

	if err != nil {
		return nil, err
	}
	var faqs []entity.FAQ
	defer rows.Close()
	for rows.Next() {
		var faq entity.FAQ
		if err := rows.Scan(&faq.Question, &faq.Answer, &faq.Topic); err != nil {
			log.Fatal(err)
		}
		faq.Topic = "stunting"
		faqs = append(faqs, faq)
	}

	return faqs, nil
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

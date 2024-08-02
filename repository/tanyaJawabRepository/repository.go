package TanyaJawabrepository

import (
	"esmartcare/entity"
)

type TanyaJawabRepository interface {
	FindAll() ([]entity.TanyaJawab, error)
	FindByValidationStatus(isValidated bool) ([]entity.TanyaJawab, error)
	FindByID(id int) (entity.TanyaJawab, error)
	Create(tanyaJawab entity.TanyaJawab) (entity.TanyaJawab, error)
	Update(tanyaJawab entity.TanyaJawab) (entity.TanyaJawab, error)
	Delete(id int) error
	FindForChatbot() ([]entity.FAQ, error)
}

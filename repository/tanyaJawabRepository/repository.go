package TanyaJawabrepository

import (
	"esmartcare/entity"
	"esmartcare/pkg/errs"
)

type TanyaJawabRepository interface {
	FindAll() ([]entity.TanyaJawab, error)
	FindByValidationStatus(isValidated bool) ([]entity.TanyaJawab, error)
	FindByID(id int) (entity.TanyaJawab, error)
	Create(tanyaJawab entity.TanyaJawab) (entity.TanyaJawab, errs.MessageErr)
	Update(tanyaJawab entity.TanyaJawab) (entity.TanyaJawab, errs.MessageErr)
	Delete(id int) error
	FindForChatbot() ([]entity.FAQ, error)
}

package pakarrepository

import (
	"esmartcare/entity"
	"esmartcare/pkg/errs"
)

type PakarRepository interface {
	CreatePakar(Pakar *entity.Pakar) (*entity.Pakar, errs.MessageErr)
	UpdatePakar(oldPakar *entity.Pakar, newPakar *entity.Pakar) (*entity.Pakar, errs.MessageErr)
	GetPakarByEmail(email string) (*entity.Pakar, errs.MessageErr)
}

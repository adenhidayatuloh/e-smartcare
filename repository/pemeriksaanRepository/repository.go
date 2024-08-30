package pemeriksaanrepository

import (
	"esmartcare/entity"
	"esmartcare/pkg/errs"
)

type PemeriksaanRepository interface {
	FindAll() ([]entity.Pemeriksaan, error)
	Create(pemeriksaan entity.Pemeriksaan) (*entity.Pemeriksaan, error)
	FindByEmail(email string, keterangan string) ([]entity.Pemeriksaan, error)
	DeleteByEmail(email string) error
	DeleteById(id int) error
	GetPemeriksaanByEmail(email string) (*entity.Pemeriksaan, error)
	UpdatePemeriksaan(oldPemeriksaan *entity.Pemeriksaan, newPemeriksaan *entity.Pemeriksaan) (*entity.Pemeriksaan, errs.MessageErr)
}

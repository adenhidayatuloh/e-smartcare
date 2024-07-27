package riwayattanyajawabrepository

import (
	"esmartcare/entity"
)

type RiwayatTanyaJawabRepository interface {
	FindByEmail(email string) ([]entity.RiwayatTanyaJawab, error)
	Create(riwayat entity.RiwayatTanyaJawab) (entity.RiwayatTanyaJawab, error)
	DeleteById(id uint) error
	FindById(id uint) (*entity.RiwayatTanyaJawab, error)
}

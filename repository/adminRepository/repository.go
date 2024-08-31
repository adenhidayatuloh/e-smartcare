package adminrepository

import (
	"esmartcare/entity"
	"esmartcare/pkg/errs"
)

type AdminRepository interface {
	CreateAdmin(Admin *entity.Admin) (*entity.Admin, errs.MessageErr)
	UpdateAdmin(oldAdmin *entity.Admin, newAdmin *entity.Admin) (*entity.Admin, errs.MessageErr)
	GetAdminByEmail(email string) (*entity.Admin, errs.MessageErr)
	
}
